package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"

	pb "crm_system/pkg/auth/authservice/gw"
)

type Server struct {
	port       string
	service    *AuthService
	grpcServer *grpc.Server
}

func NewServer(
	port string,
	service *AuthService,
) *Server {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	return &Server{
		port:       port,
		service:    service,
		grpcServer: grpcServer,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return fmt.Errorf("failed to listen grpc port: %s", s.port)
	}

	pb.RegisterAuthServiceServer(s.grpcServer, s.service)

	go s.grpcServer.Serve(listener)

	return nil
}

func (s *Server) Close() {
	s.grpcServer.GracefulStop()
}
