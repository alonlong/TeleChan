package api

import (
	"TeleChan/pkg/config"
	"TeleChan/pkg/log"
	"net"
	"sync"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	// DefaultKeepaliveMinTime - if a client pings more than once every 5 seconds, terminate the connection
	DefaultKeepaliveMinTime = 5
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             DefaultKeepaliveMinTime * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,                                  // Allow pings even when there are no active streams
}

// Service represents the interfaces for TeleChan service
type Service struct {
	server   *grpc.Server   // the grpc server
	executor *apiExecutor   // the api methods executor
	settings *config.Config // settings for the server
}

// NewService returns an internal implemented TeleChan service
func NewService(settings *config.Config) *Service {
	s := &Service{
		settings: settings,
	}

	// init the api executor
	s.executor = newAPIExecutor(settings)

	// init the grpc server
	s.server = grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			auth.StreamServerInterceptor(s.executor.auth),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			auth.UnaryServerInterceptor(s.executor.auth),
		)),
	)
	// register the TeleChan service
	RegisterTeleChanServer(s.server, s.executor)

	return s
}

// Start the grpc server
func (s *Service) Start(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", s.settings.Server.ListenAddr)
	if err != nil {
		panic(err)
	}

	// start a goroutine for grpc serve
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.server.Serve(lis); err != nil {
			log.Infof("grpc serve: %v\n", err)
		}
	}()

	// start the goroutines for pushing messages
	s.executor.startPush(wg)
	// check if the goroutines are ready
	s.executor.isReady()
}

// Stop releases the GRPC server's resources
func (s *Service) Stop() {
	if s.executor != nil {
		s.executor.close()
	}
	if s.server != nil {
		s.server.GracefulStop()
	}
}
