package app

import (
	"crm_system/config/auth"
	"crm_system/internal/auth/controller/grpc"
	middleware2 "crm_system/internal/auth/controller/http/middleware"
	"crm_system/internal/auth/controller/http/v1"
	repoPkg "crm_system/internal/auth/repository"
	servicePkg "crm_system/internal/auth/service"
	httpserver2 "crm_system/pkg/auth/httpserver"
	"crm_system/pkg/auth/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *auth.Configuration) {
	l := logger.New(cfg.Gin.Mode)
	repo := repoPkg.New(cfg, l)

	// Migrate the tables with gorm.Migrator
	Migrate(repo, l)

	service := servicePkg.New(cfg, repo, l)
	middleware := middleware2.New(repo, cfg)
	handler := gin.Default()

	grpcService := grpc.NewService(l, repo, cfg)
	grpcServer := grpc.NewServer(cfg.Grpc.Port, grpcService)
	err := grpcServer.Start()
	if err != nil {
		log.Panicf("failed to start grpc-server err: %v", err)
	}

	defer grpcServer.Close()
	v1.NewRouter(handler, service, middleware)
	httpServer := httpserver2.New(handler, cfg, httpserver2.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("auth - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("auth - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("auth - Run - httpServer.Shutdown: %w", err))
	}

}
