package api

import (
	"TeleChan/pkg/config"
	"TeleChan/pkg/db"
)

// Service represents the interfaces for TeleChan service
type Service struct {
	UnimplementedTeleChanServer

	settings *config.Config // settings for the server
	handler  *db.Handler    // handler for db operation
	closed   chan struct{}  // receive a close signal
}

// NewService returns an internal implemented TeleChan service
func NewService(settings *config.Config) *Service {
	s := &Service{
		settings: settings,
	}

	// init the database handler
	s.handler = db.NewHandler(settings)

	return s
}

// Close releases the GRPC service's resources
func (s *Service) Close() {
	if s.closed != nil {
		close(s.closed)
	}
	if s.handler != nil {
		s.handler.Close()
	}
}
