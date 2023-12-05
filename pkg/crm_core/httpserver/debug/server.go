package debug

import (
	"context"
	"crm_system/config/crm_core"
	"log"
	"net/http"
	"time"
)

// Server -.
type ServerDebug struct {
	server          *http.Server
	config          *crm_core.Configuration
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler, config *crm_core.Configuration, opts ...Option) *ServerDebug {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  time.Duration(config.HTTP.DefaultReadTimeout),
		WriteTimeout: time.Duration(config.HTTP.DefaultWriteTimeout),
		Addr:         config.HTTP.DebugPort,
	}

	s := &ServerDebug{
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

func (s *ServerDebug) start() {
	go func() {
		log.Printf("Listening debug server on port:%s", s.config.HTTP.DebugPort)
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *ServerDebug) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *ServerDebug) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
