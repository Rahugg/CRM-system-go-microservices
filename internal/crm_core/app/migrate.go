package app

import (
	entityRepo "crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/repository"
	"crm_system/pkg/crm_core/logger"
	"fmt"
)

func Migrate(repo *repository.CRMSystemRepo, l *logger.Logger) {
	repo.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := repo.DB.AutoMigrate(
		&entityRepo.Company{},
		&entityRepo.Contact{},
		&entityRepo.Deal{},
		&entityRepo.Task{},
		&entityRepo.Task{},
		&entityRepo.Role{},
	)
	if err != nil {
		l.Fatal("Automigration failed")
	}

	fmt.Println("👍 Migration complete")
}
