// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"crm_system/config/auth"
	"log"
	"net/http"
	"time"
)

// Server -.
type Server struct {
	server          *http.Server
	config          *auth.Configuration
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler, config *auth.Configuration, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  time.Duration(config.HTTP.DefaultReadTimeout),
		WriteTimeout: time.Duration(config.HTTP.DefaultWriteTimeout),
		Addr:         config.HTTP.Port,
	}

	s := &Server{
		server:          httpServer,
		config:          config,
		notify:          make(chan error, 1),
		shutdownTimeout: time.Duration(config.HTTP.DefaultShutdownTimeout),
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		log.Printf("Listening on port:%s", s.config.HTTP.Port)
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
