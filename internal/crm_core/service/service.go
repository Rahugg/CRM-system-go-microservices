package service

import (
	"crm_system/config/crm_core"
	"crm_system/internal/crm_core/repository"
	"crm_system/pkg/crm_core/logger"
)

type Service struct {
	Repo   *repository.CRMSystemRepo
	Config *crm_core.Configuration
}

func New(config *crm_core.Configuration, repo *repository.CRMSystemRepo, l *logger.Logger) *Service {
	return &Service{
		Repo:   repo,
		Config: config,
	}
}
