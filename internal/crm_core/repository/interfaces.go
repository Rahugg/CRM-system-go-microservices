package repository

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

type (
	CompanyRepo interface {
		GetCompanies(ctx *gin.Context) (*[]entity.Company, error)
		GetCompany(ctx *gin.Context, id string) (*entity.Company, error)
		CreateCompany(ctx *gin.Context, company *entity.Company) error
		SaveCompany(ctx *gin.Context, newCompany entity.Company) error
		DeleteCompany(ctx *gin.Context, id string) error
	}

	TaskRepo interface {
		GetTasks(ctx *gin.Context) (*[]entity.Task, error)
		GetTask(ctx *gin.Context, id string) (*entity.Task, error)
		CreateTask(ctx *gin.Context, task *entity.Task) error
		SaveTask(ctx *gin.Context, newTask *entity.Task) error
		DeleteTask(ctx *gin.Context, id string, task *entity.Task) error
	}

	ContactRepo interface {
		GetContacts(ctx *gin.Context) (*[]entity.Contact, error)
		GetContact(ctx *gin.Context, id string) (*entity.Contact, error)
		CreateContact(ctx *gin.Context, contact *entity.Contact) error
		SaveContact(ctx *gin.Context, newContact *entity.Contact) error
		DeleteContact(ctx *gin.Context, id string, contact *entity.Contact) error
	}
	DealRepo interface {
		GetDeals(ctx *gin.Context) (*[]entity.Deal, error)
		GetDeal(ctx *gin.Context, id string) (*entity.Deal, error)
		CreateDeal(ctx *gin.Context, deal *entity.Deal) error
		SaveDeal(ctx *gin.Context, newDeal *entity.Deal) error
		DeleteDeal(ctx *gin.Context, id string, deal *entity.Deal) error
	}

	TicketRepo interface {
		GetTickets(ctx *gin.Context) (*[]entity.Ticket, error)
		GetTicket(ctx *gin.Context, id string) (*entity.Ticket, error)
		CreateTicket(ctx *gin.Context, ticket *entity.Ticket) error
		SaveTicket(ctx *gin.Context, newTicket *entity.Ticket) error
		DeleteTicket(ctx *gin.Context, id string, deal *entity.Ticket) error
	}
)
