package service

import (
	"crm_system/config/auth"
	"crm_system/internal/auth/repository"
	"crm_system/pkg/auth/logger"
)

type Service struct {
	Repo   *repository.AuthRepo
	Config *auth.Configuration
}

func New(config *auth.Configuration, repo *repository.AuthRepo, l *logger.Logger) *Service {
	return &Service{
		Repo:   repo,
		Config: config,
	}
}
