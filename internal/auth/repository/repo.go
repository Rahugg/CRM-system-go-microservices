package repository

import (
	"crm_system/config/auth"
	"crm_system/pkg/auth/logger"
	"crm_system/pkg/auth/postgres"
	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB
}

func New(config *auth.Configuration, l *logger.Logger) *AuthRepo {
	db := postgres.ConnectDB(config, l)
	return &AuthRepo{
		DB: db,
	}
}
