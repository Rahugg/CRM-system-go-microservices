package main

import (
	"crm_system/config/crm_core"
	entityRepo "crm_system/internal/crm_core/entity"
	repoPkg "crm_system/internal/crm_core/repository"
	"crm_system/pkg/crm_core/logger"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	cfg := crm_core.NewConfig()
	l := logger.New(cfg.Gin.Mode)
	repo := repoPkg.New(cfg, l)

	company := &entityRepo.Company{
		Name:      "Sample Company",
		Address:   "123 Main Street",
		Phone:     "555-1234",
		ManagerID: uuid.New(),
	}
	if repo.DB.Model(&company).Where("phone = ?", company.Phone).Updates(&company).RowsAffected == 0 {
		repo.DB.Create(&company)
	}

	contact := &entityRepo.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "555-5678",
		CompanyID: 1, // Assuming a valid CompanyID from your Company data
	}
	if repo.DB.Model(&contact).Where("company_id= ?", contact.CompanyID).Updates(&contact).RowsAffected == 0 {
		repo.DB.Create(&contact)
	}

	deal := &entityRepo.Deal{
		Title:     "Sample Deal",
		Value:     100000,
		Status:    entityRepo.InProgress,
		ContactID: 1,
	}
	if repo.DB.Model(&deal).Where("contact_id= ?", deal.ContactID).Updates(&deal).RowsAffected == 0 {
		repo.DB.Create(&deal)
	}

	task := &entityRepo.Task{
		Name:             "Sample Task",
		Description:      "This is a sample task description.",
		DueDate:          time.Now().AddDate(0, 0, 7),
		AssignedTo:       uuid.New(),
		AssociatedDealID: 1,
		State:            "Open",
	}
	if repo.DB.Model(&task).Where("associated_deal_id= ?", task.AssociatedDealID).Updates(&task).RowsAffected == 0 {
		repo.DB.Create(&task)
	}

	vote := &entityRepo.Vote{
		SenderID: uuid.New(),
		TaskID:   1,
	}

	if repo.DB.Model(&vote).Where("task_id= ?", vote.TaskID).Updates(&vote).RowsAffected == 0 {
		repo.DB.Create(&vote)
	}

	ticket := &entityRepo.Ticket{
		IssueDescription: "Sample ticket issue description.",
		Status:           entityRepo.Open,
		ContactID:        1,          // Assuming a valid ContactID from your Contact data
		AssignedTo:       uuid.New(), // Generating a random UUID for AssignedTo
	}
	if repo.DB.Model(&ticket).Where("contact_id= ?", ticket.ContactID).Updates(&ticket).RowsAffected == 0 {
		repo.DB.Create(&ticket)
	}

	fmt.Println("Mock data inserted successfully✅ crm_core-service")
}
