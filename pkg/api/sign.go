package api

import (
	"TeleChan/pkg/db/model"
	"TeleChan/pkg/log"
	context "context"
	"time"

	"github.com/dgrijalva/jwt-go"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

var (
	// jwt secret for generating the jwt token
	jwtSecret = []byte("pngOEtu35ewyKJWZ")
)

// Create a struct that will be encoded to a JWT
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Refresh - refresh a new token with the current token
func (e *apiExecutor) Refresh(ctx context.Context, request *RefreshRequest) (*RefreshReply, error) {
	// fetch the user id from context
	userID := e.getUserId(ctx)

	// declare the expiration time of the token
	expiration := time.Now().Add(e.settings.Server.TokenExpire * time.Minute)
	// create the jwt claims, which include the username and expiration time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	// decleare the token with the algorithem used for signing
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Generate token: %v", err)
	}

	return &RefreshReply{RefreshToken: token}, nil
}

// SignUp - signup with user's email and password, need to check if it's existed in database
func (e *apiExecutor) SignUp(ctx context.Context, request *SignUpRequest) (*SuccessReply, error) {
	if request.User.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email is empty")
	}
	if request.User.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Password is empty")
	}

	// query the user from database by email
	u, err := e.handler.FindUserByEmail(request.User.Email)
	if err != nil {
		log.Errorf("Find user by email %s: %v", request.User.Email, err)
		return nil, status.Errorf(codes.Internal, "Find user: %v", err)
	}
	if u != nil {
		return nil, status.Errorf(codes.AlreadyExists, "User is existed")
	}

	u = &model.User{
		UserID:   newUUID(),
		Email:    request.User.Email,
		Password: request.User.Password,
	}
	// insert the user to database
	if err := e.handler.CreateUser(u); err != nil {
		log.Errorf("Create user: %v", err)
		return nil, status.Errorf(codes.Internal, "Create user: %v", err)
	}

	return &SuccessReply{}, nil
}

// SignIn - signin with user's email and password, after success, return a jwt token
func (e *apiExecutor) SignIn(ctx context.Context, request *SignInRequest) (*SignInReply, error) {
	if request.User.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email is empty")
	}
	if request.User.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Password is empty")
	}

	// query the user from database by email
	u, err := e.handler.FindUserByEmail(request.User.Email)
	if err != nil {
		log.Errorf("Find user by email %s: %v", request.User.Email, err)
		return nil, status.Errorf(codes.Internal, "Find user: %v", err)
	}
	if u == nil {
		return nil, status.Errorf(codes.NotFound, "User is not found")
	}

	// if the password is right
	if u.Password != request.User.Password {
		return nil, status.Errorf(codes.Unauthenticated, "Password is wrong")
	}

	// declare the expiration time of the token
	expiration := time.Now().Add(e.settings.Server.TokenExpire * time.Minute)
	// create the jwt claims, which include the username and expiration time
	claims := &Claims{
		UserID: u.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	// decleare the token with the algorithem used for signing
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Generate token: %v", err)
	}

	return &SignInReply{Token: token}, nil
}
