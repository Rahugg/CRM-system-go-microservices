package postgres

import (
	"crm_system/config/crm_core"
	"crm_system/pkg/crm_core/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *crm_core.Configuration, l *logger.Logger) *gorm.DB {
	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
	)
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		l.Fatal(err)
	}
	return db
}
