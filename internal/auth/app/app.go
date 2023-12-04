package app

import (
	"crm_system/config/auth"
	"crm_system/internal/auth/controller/consumer"
	"crm_system/internal/auth/controller/grpc"
	middleware2 "crm_system/internal/auth/controller/http/middleware"
	"crm_system/internal/auth/controller/http/v1"
	_ "crm_system/internal/auth/docs"
	repoPkg "crm_system/internal/auth/repository"
	servicePkg "crm_system/internal/auth/service"
	storagePkg "crm_system/internal/auth/storage"
	"crm_system/internal/kafka"
	httpserver2 "crm_system/pkg/auth/httpserver"
	"crm_system/pkg/auth/logger"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *auth.Configuration) {
	l := logger.New(cfg.Gin.Mode)
	repo := repoPkg.New(cfg, l)
	userVerificationProducer, err := kafka.NewProducer(cfg)
	if err != nil {
		l.Fatal("failed NewProducer err: %v", err)
	}

	userVerificationConsumerCallback := consumer.NewUserVerificationCallback(l)

	userVerificationConsumer, err := kafka.NewConsumer(l, cfg, userVerificationConsumerCallback)
	if err != nil {
		l.Fatal("failed NewConsumer err: %v", err)
	}

	go userVerificationConsumer.Start()

	storage := storagePkg.NewDataStorage(cfg.Storage.Interval, repo, l)
	service := servicePkg.New(cfg, repo, userVerificationProducer, storage)

	go storage.Run()

	middleware := middleware2.New(repo, cfg)
	handler := gin.Default()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	grpcService := grpc.NewService(l, repo, cfg)
	grpcServer := grpc.NewServer(cfg.Grpc.Port, grpcService)
	err = grpcServer.Start()
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
