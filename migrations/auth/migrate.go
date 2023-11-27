package main

import (
	"crm_system/config/auth"
	entityRepo "crm_system/internal/auth/entity"
	_ "crm_system/internal/auth/repository"
	repoPkg "crm_system/internal/auth/repository"
	"crm_system/pkg/auth/logger"
	"fmt"
)

func main() {
	cfg := auth.NewConfig()
	l := logger.New(cfg.Gin.Mode)
	repo := repoPkg.New(cfg, l)
	repo.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := repo.DB.AutoMigrate(
		&entityRepo.User{},
		&entityRepo.Role{},
		&entityRepo.UserCode{},
	)
	if err != nil {
		l.Fatal("Automigration failed")
	}

	fmt.Println("👍 Migration complete - auth service")
}
