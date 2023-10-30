package postgres

import (
	"crm_system/config/auth"
	"crm_system/pkg/auth/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *auth.Configuration, l *logger.Logger) *gorm.DB {
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
