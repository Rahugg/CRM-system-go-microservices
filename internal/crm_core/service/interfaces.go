package service

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

type (
	CompanyService interface {
		GetCompanies(ctx *gin.Context) (*[]entity.Company, error)
		GetCompany(ctx *gin.Context, id string) (*entity.Company, error)
		CreateCompany(ctx *gin.Context, company entity.Company) error
		UpdateCompany(ctx *gin.Context, newCompany entity.NewCompany, id string) error
		DeleteCompany(ctx *gin.Context, id string) error
	}

	TaskService interface {
		GetTasks(ctx *gin.Context) (*[]entity.Task, error)
		GetTask(ctx *gin.Context, id string) (*entity.Task, error)
		CreateTask(ctx *gin.Context, task entity.Task) error
		UpdateTask(ctx *gin.Context, newTask entity.Task, id string) error
		DeleteTask(ctx *gin.Context, id string) error
	}

	ContactService interface {
		GetContacts(ctx *gin.Context) (*[]entity.Contact, error)
		GetContact(ctx *gin.Context, id string) (*entity.Contact, error)
		CreateContact(ctx *gin.Context, contact entity.Contact) error
		UpdateContact(ctx *gin.Context, newContact entity.Contact, id string) error
		DeleteContact(ctx *gin.Context, id string) error
	}
)
