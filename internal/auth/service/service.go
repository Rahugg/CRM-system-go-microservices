package service

import (
	"crm_system/config/auth"
	"crm_system/internal/auth/repository"
	"crm_system/internal/kafka"
)

type Service struct {
	Repo                     *repository.AuthRepo
	Config                   *auth.Configuration
	userVerificationProducer *kafka.Producer
}

func New(config *auth.Configuration, repo *repository.AuthRepo, userVerificationProducer *kafka.Producer) *Service {
	return &Service{
		Repo:                     repo,
		Config:                   config,
		userVerificationProducer: userVerificationProducer,
	}
}
