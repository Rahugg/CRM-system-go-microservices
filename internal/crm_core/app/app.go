package app

import (
	"crm_system/config/crm_core"
	middleware2 "crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/controller/http/v1"
	repoPkg "crm_system/internal/crm_core/repository"
	servicePkg "crm_system/internal/crm_core/service"
	"crm_system/internal/crm_core/transport"
	"crm_system/pkg/crm_core/cache"
	httpserverPkg "crm_system/pkg/crm_core/httpserver"
	"crm_system/pkg/crm_core/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *crm_core.Configuration) {
	l := logger.New(cfg.Gin.Mode)
	repo := repoPkg.New(cfg, l)
	// migrate the tables with gorm.Migrator
	Migrate(repo, l)

	//REDIS implementation
	redisClient, err := cache.NewRedisClient()
	if err != nil {
		return
	}

	contactCache := cache.NewContactCache(redisClient, 10*time.Minute)

	validateGrpcTransport := transport.NewValidateGrpcTransport(*cfg)

	service := servicePkg.New(cfg, repo, l)
	middleware := middleware2.New(repo, cfg, validateGrpcTransport)
	handler := gin.Default()

	v1.NewRouter(handler, service, middleware, contactCache)
	httpServer := httpserverPkg.New(handler, cfg, httpserverPkg.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("crm_system - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("crm_system - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("crm_system - Run - httpServer.Shutdown: %w", err))
	}

}
