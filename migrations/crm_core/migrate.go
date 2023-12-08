package main

import (
	"crm_system/config/crm_core"
	entityRepo "crm_system/internal/crm_core/entity"
	repoPkg "crm_system/internal/crm_core/repository"
	"crm_system/pkg/crm_core/logger"
	"fmt"
)

func main() {
	cfg := crm_core.NewConfig()
	l := logger.New(cfg.Gin.Mode)
	repo := repoPkg.New(cfg, l)
	repo.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := repo.DB.AutoMigrate(
		&entityRepo.Company{},
		&entityRepo.Contact{},
		&entityRepo.Deal{},
		&entityRepo.Task{},
		&entityRepo.Ticket{},
		&entityRepo.Vote{},
		&entityRepo.TaskChanges{},
	)
	if err != nil {
		l.Fatal("Automigration failed")
	}

	fmt.Println("👍 Migration complete - crm_core service")
}
