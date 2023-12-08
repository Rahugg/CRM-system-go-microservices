package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "crm_system/pkg/auth/authservice/gw"
)

type Server struct {
	port       string
	service    *Service
	grpcServer *grpc.Server
}

func NewServer(
	port string,
	service *Service,
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

	go func() {
		err = s.grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("error happened in grpc: %s", err)
		}
	}()

	return nil
}

func (s *Server) Close() {
	s.grpcServer.GracefulStop()
}
