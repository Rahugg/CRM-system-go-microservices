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
		GetTasks(ctx *gin.Context, dealId string, user entity.User) ([]map[string]interface{}, error)
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

	TicketService interface {
		GetTickets(ctx *gin.Context) (*[]entity.Ticket, error)
		GetTicket(ctx *gin.Context, id string) (*entity.Ticket, error)
		CreateTicket(ctx *gin.Context, ticket entity.Ticket) error
		UpdateTicket(ctx *gin.Context, newTicket entity.Ticket, id string) error
		DeleteTicket(ctx *gin.Context, id string) error
	}

	DealService interface {
		GetDeals(ctx *gin.Context) (*[]entity.Deal, error)
		GetDeal(ctx *gin.Context, id string) (*entity.Deal, error)
		CreateDeal(ctx *gin.Context, deal entity.Deal) error
		UpdateDeal(ctx *gin.Context, newDeal entity.Deal, id string) error
		DeleteDeal(ctx *gin.Context, id string) error
	}
)
