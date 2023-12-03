package service

import (
	"crm_system/config/auth"
	"crm_system/internal/auth/repository"
	"crm_system/internal/auth/storage"
	"crm_system/internal/kafka"
)

type Service struct {
	Repo                     *repository.AuthRepo
	Config                   *auth.Configuration
	userVerificationProducer *kafka.Producer
	Storage                  *storage.DataStorage
}

func New(config *auth.Configuration, repo *repository.AuthRepo, userVerificationProducer *kafka.Producer, storage *storage.DataStorage) *Service {
	return &Service{
		Repo:                     repo,
		Config:                   config,
		userVerificationProducer: userVerificationProducer,
		Storage:                  storage,
	}
}
