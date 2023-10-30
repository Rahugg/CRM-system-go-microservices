package app

import (
	entityRepo "crm_system/internal/auth/entity"
	"crm_system/internal/auth/repository"
	"crm_system/pkg/auth/logger"
	"fmt"
)

func Migrate(repo *repository.AuthRepo, l *logger.Logger) {
	repo.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := repo.DB.AutoMigrate(
		&entityRepo.User{},
		&entityRepo.Role{},
	)
	if err != nil {
		l.Fatal("Automigration failed")
	}
	roles := []*entityRepo.Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "manager"},
		{ID: 3, Name: "sales_rep"},
		{ID: 4, Name: "support_rep"},
		{ID: 5, Name: "guest"},
	}
	for i, role := range roles {
		if repo.DB.Model(&role).Where("id = ?", i+1).Updates(&role).RowsAffected == 0 {
			repo.DB.Create(&role)
		}
	}

	newAdmin := entityRepo.User{
		FirstName: "admin",
		LastName:  "main",
		Email:     "a_faizolla@kbtu.kz",
		Password:  "$2a$12$/84UJA1OqAVDl.6BB9r5VegjczNvNXM.DlaFYF8uk9QoB6YK2LdoK",
		RoleID:    1,
		Provider:  "admin",
	}

	if repo.DB.Model(&newAdmin).Where("email = ?", "a_faizolla@kbtu.kz").Updates(&newAdmin).RowsAffected == 0 {
		repo.DB.Create(&newAdmin)
	}

	fmt.Println("👍 Migration complete")
}
