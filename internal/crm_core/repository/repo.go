package repository

import (
	"crm_system/config/crm_core"
	"crm_system/pkg/crm_core/logger"
	"crm_system/pkg/crm_core/postgres"
	"gorm.io/gorm"
)

type CRMSystemRepo struct {
	DB *gorm.DB
}

func New(config *crm_core.Configuration, l *logger.Logger) *CRMSystemRepo {
	db := postgres.ConnectDB(config, l)
	return &CRMSystemRepo{
		DB: db,
	}
}
