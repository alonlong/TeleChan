package api

import (
	"TeleChan/pkg/config"
	"TeleChan/pkg/db"
	"TeleChan/pkg/log"
	"context"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
)

var (
	// DefaultPassMethods - these methods don't need to authorize
	DefaultPassMethods = map[string]bool{
		"/api.TeleChan/SignUp": true,
		"/api.TeleChan/SignIn": true,
	}
)

// apiExecutor can run the GRPC API methods
type apiExecutor struct {
	UnimplementedTeleChanServer

	settings *config.Config         // settings for the server
	handler  *db.Handler            // handler for db operation
	closed   chan struct{}          // receive a close signal
	users    map[string]*Connection // user's connections
	userMux  sync.RWMutex           // mutex for users
	ready    chan struct{}          // mark the push goroutines ready
	messages chan []*Carrier        // queue for pushing messages
}

// newAPIExecutor returns the api methods executor
func newAPIExecutor(settings *config.Config) *apiExecutor {
	e := &apiExecutor{
		users:    make(map[string]*Connection),
		closed:   make(chan struct{}),
		settings: settings,
	}

	// init the ready channel for starting goroutines
	e.ready = make(chan struct{}, e.settings.Server.PushRoutines)

	// init the queue for pushing messages
	e.messages = make(chan []*Carrier, e.settings.Server.QueueSize)

	// init the database handler
	e.handler = db.NewHandler(settings)

	return e
}

// authorize the request's access permissions with token
func (e *apiExecutor) auth(ctx context.Context) (context.Context, error) {
	// resolve the method name
	method, ok := grpc.Method(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "Invalid grpc method")
	}

	// if the method can be passed
	if _, ok := DefaultPassMethods[method]; ok {
		return ctx, nil
	}

	// resolve the token from context
	jwtToken := e.getJwtToken(ctx)
	if jwtToken == "" {
		return ctx, status.Errorf(codes.Unauthenticated, "Token is empty")
	}

	// initialize a new instance of 'Claims'
	claims := &Claims{}
	// parse the jwt string into claims
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	// check if the token is valid or if the signature is matched
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "Parse token: %v", err)
	}
	if token.Valid == false {
		return ctx, status.Errorf(codes.Unauthenticated, "Token is invalid")
	}

	now := time.Now()
	// check if the token is expired
	if claims.ExpiresAt-now.Unix() < 30 {
		return ctx, status.Errorf(codes.Unauthenticated, "Token is expired")
	}

	// append the 'user-id' to context
	newCtx := metadata.NewIncomingContext(ctx, metadata.Pairs(headerUser, claims.UserID))

	return newCtx, nil
}

// start the goroutines for pushing messages
func (e *apiExecutor) startPush(wg *sync.WaitGroup) {
	// start a goroutin pool for pushing messages
	for i := 0; i < e.settings.Server.PushRoutines; i++ {
		wg.Add(1)
		go e.handlePush(wg)
	}
}

// isReady check if the push goroutines are ready
func (e *apiExecutor) isReady() {
	for i := 0; i < e.settings.Server.PushRoutines; i++ {
		<-e.ready
	}
	log.Info("push goroutines are ready")
}

// close the database connections and other resources
func (e *apiExecutor) close() {
	if e.closed != nil {
		close(e.closed)
	}
	if e.handler != nil {
		e.handler.Close()
	}
}
