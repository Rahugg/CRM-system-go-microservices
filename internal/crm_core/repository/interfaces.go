package repository

import (
	"crm_system/internal/crm_core/entity"
)

type (
	CompanyRepo interface {
		GetCompanies() (*[]entity.Company, error)
		GetCompany(id string) (*entity.Company, error)
		CreateCompany(company *entity.Company) error
		SaveCompany(newCompany *entity.Company) error
		DeleteCompany(id string, company *entity.Company) error
		SearchCompany(query string) (*[]entity.Company, error)
		SortCompanies(companies *[]entity.Company, sortBy, sortOrder string) (*[]entity.Company, error)
		FilterCompaniesByPhone(companies *[]entity.Company, phone string) (*[]entity.Company, error)
	}
	ContactRepo interface {
		GetContacts() (*[]entity.Contact, error)
		GetContact(id string) (*entity.Contact, error)
		CreateContact(contact *entity.Contact) error
		SaveContact(newContact *entity.Contact) error
		DeleteContact(id string, contact *entity.Contact) error
		SearchContact(query string) (*[]entity.Contact, error)
		SortContacts(contacts *[]entity.Contact, sortBy, sortOrder string) (*[]entity.Contact, error)
		FilterContactsByPhone(contacts *[]entity.Contact, phone string) (*[]entity.Contact, error)
	}

	DealRepo interface {
		GetDeals() (*[]entity.Deal, error)
		GetDeal(id string) (*entity.Deal, error)
		CreateDeal(deal *entity.Deal) error
		SaveDeal(newDeal *entity.Deal) error
		DeleteDeal(id string, deal *entity.Deal) error
		SearchDeal(query string) (*[]entity.Deal, error)
		SortDeals(deals *[]entity.Deal, sortBy, sortOrder string) (*[]entity.Deal, error)
		FilterDealsByStatus(deals *[]entity.Deal, status string) (*[]entity.Deal, error)
	}

	TaskRepo interface {
		GetTasks() (*[]entity.Task, error)
		GetTask(id string) (*entity.Task, error)
		GetTasksByDealId(dealId string) ([]entity.Task, error)
		CreateTaskChanges(taskChanges *entity.TaskChanges) error
		CreateTask(newTask *entity.TaskInput) error
		CreateVote(user *entity.User, voteInput *entity.VoteInput) error
		GetChangesOfTask(id string) (*[]entity.TaskChanges, error)
		SaveTask(newTask *entity.Task) error
		DeleteTask(id string, task *entity.Task) error
		SearchTask(query string) (*[]entity.Task, error)
		FilterTasksByStates(tasks []map[string]interface{}, state string) ([]map[string]interface{}, error)
	}

	TicketRepo interface {
		GetTickets() (*[]entity.Ticket, error)
		GetTicket(id string) (*entity.Ticket, error)
		CreateTicket(ticket *entity.Ticket) error
		SaveTicket(newTicket *entity.Ticket) error
		DeleteTicket(id string, deal *entity.Ticket) error
		SearchTicket(query string) (*[]entity.Ticket, error)
		SortTickets(tickets *[]entity.Ticket, sortBy, sortOrder string) (*[]entity.Ticket, error)
		FilterTicketsByStatus(tickets *[]entity.Ticket, status string) (*[]entity.Ticket, error)
	}
)
