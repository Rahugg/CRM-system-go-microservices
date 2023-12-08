package main

import (
	"crm_system/config/crm_core"
	entityRepo "crm_system/internal/crm_core/entity"
	repoPkg "crm_system/internal/crm_core/repository"
	"crm_system/pkg/crm_core/logger"
)

func main() {
	cfg := crm_core.NewConfig()
	l := logger.New(cfg.Gin.Mode)
	repo := repoPkg.New(cfg, l)

	err := repo.DB.Migrator().DropTable(&entityRepo.Company{},
		&entityRepo.Contact{},
		&entityRepo.Deal{},
		&entityRepo.Task{},
		&entityRepo.Ticket{},
		&entityRepo.Vote{},
		&entityRepo.TaskChanges{})
	if err != nil {
		l.Error("error happened: %s", err)
		return
	}
}
