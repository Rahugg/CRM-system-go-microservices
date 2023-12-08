package debug

import (
	"net"
	"time"
)

// Option -.
type Option func(*ServerDebug)

// Port -.
func Port(port string) Option {
	return func(s *ServerDebug) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *ServerDebug) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *ServerDebug) {
		s.server.WriteTimeout = timeout
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *ServerDebug) {
		s.shutdownTimeout = timeout
	}
}
