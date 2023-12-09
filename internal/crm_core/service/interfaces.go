package service

import (
	"crm_system/internal/crm_core/entity"
)

type (
	CompanyService interface {
		GetCompanies(sortBy, sortOrder, phone string) (*[]entity.Company, error)
		GetCompany(id string) (*entity.Company, error)
		CreateCompany(company entity.Company) error
		UpdateCompany(newCompany entity.NewCompany, id string) error
		DeleteCompany(id string) error
		SearchCompany(query string) (*[]entity.Company, error)
		sortCompanies(companies *[]entity.Company, sortBy, sortOrder string) (*[]entity.Company, error)
		filterCompaniesByPhone(companies *[]entity.Company, phone string) (*[]entity.Company, error)
	}

	TaskService interface {
		GetTasks(dealId, sortBy, sortOrder, stateInput string, user *entity.User) ([]map[string]interface{}, error)
		GetTask(id string) (*entity.Task, error)
		CreateTask(task *entity.TaskInput) error
		UpdateTask(newTask entity.TaskEditInput, id string, user *entity.User) error
		GetChangesOfTask(id string) (*[]entity.TaskChanges, error)
		DeleteTask(id string) error
		SearchTask(query string) (*[]entity.Task, error)
		sortTasks(columns []map[string]interface{}, sortBy, sortOrder, state string) ([]map[string]interface{}, error)
		filterTasksByStates(columns []map[string]interface{}, state string) (map[string]interface{}, error)
		Vote(user *entity.User, voteInput *entity.VoteInput) error
	}

	ContactService interface {
		GetContacts(sortBy, sortOrder, phone string) (*[]entity.Contact, error)
		GetContact(id string) (*entity.Contact, error)
		CreateContact(contact entity.Contact) error
		UpdateContact(newContact entity.Contact, id string) error
		DeleteContact(id string) error
		SearchContact(query string) (*[]entity.Contact, error)
		sortContacts(contacts *[]entity.Contact, sortBy, sortOrder string) (*[]entity.Contact, error)
		filterContactsByPhone(contacts *[]entity.Contact, phone string) (*[]entity.Contact, error)
	}

	TicketService interface {
		GetTickets(sortBy, sortOrder, status string) (*[]entity.Ticket, error)
		GetTicket(id string) (*entity.Ticket, error)
		CreateTicket(ticket entity.Ticket) error
		UpdateTicket(newTicket entity.Ticket, id string) error
		DeleteTicket(id string) error
		SearchTicket(query string) (*[]entity.Ticket, error)
		sortTickets(tickets *[]entity.Ticket, sortBy, sortOrder string) (*[]entity.Ticket, error)
		filterTicketsByStatus(tickets *[]entity.Ticket, status string) (*[]entity.Ticket, error)
	}

	DealService interface {
		GetDeals(sortBy, sortOrder, status string) (*[]entity.Deal, error)
		GetDeal(id string) (*entity.Deal, error)
		CreateDeal(deal entity.Deal) error
		UpdateDeal(newDeal entity.Deal, id string) error
		DeleteDeal(id string) error
		SearchDeal(query string) (*[]entity.Deal, error)
		sortDeals(deals *[]entity.Deal, sortBy, sortOrder string) (*[]entity.Deal, error)
		filterDealsByStatus(deals *[]entity.Deal, status string) (*[]entity.Deal, error)
	}
)
