package service

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) GetTickets(ctx *gin.Context, sortBy, sortOrder, status string) (*[]entity.Ticket, error) {
	tickets, err := s.Repo.GetTickets(ctx)
	if status != "" {
		tickets, err = s.filterTicketsByStatus(tickets, status)
	}

	if sortBy != "" {
		tickets, err = s.sortTickets(tickets, sortBy, sortOrder)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *Service) sortTickets(tickets *[]entity.Ticket, sortBy, sortOrder string) (*[]entity.Ticket, error) {
	tickets, err := s.Repo.SortTickets(tickets, sortBy, sortOrder)

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *Service) filterTicketsByStatus(tickets *[]entity.Ticket, status string) (*[]entity.Ticket, error) {
	tickets, err := s.Repo.FilterTicketsByStatus(tickets, status)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *Service) GetTicket(ctx *gin.Context, id string) (*entity.Ticket, error) {
	ticket, err := s.Repo.GetTicket(ctx, id)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}
func (s *Service) CreateTicket(ctx *gin.Context, ticket entity.Ticket) error {
	if err := s.Repo.CreateTicket(ctx, &ticket); err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateTicket(ctx *gin.Context, newTicket entity.Ticket, id string) error {
	ticket, err := s.Repo.GetTicket(ctx, id)
	if err != nil {
		return err
	}

	if newTicket.IssueDescription != "" {
		ticket.IssueDescription = newTicket.IssueDescription
	}

	if isValidTicketStatus(newTicket.Status) {
		ticket.Status = newTicket.Status
	}

	if newTicket.ContactID != 0 {
		ticket.ContactID = newTicket.ContactID
	}

	if newTicket.AssignedTo != uuid.Nil {
		ticket.AssignedTo = newTicket.AssignedTo
	}

	if err = s.Repo.SaveTicket(ctx, ticket); err != nil {
		return err
	}

	return nil
}

func isValidTicketStatus(status entity.StatusTicket) bool {
	switch status {
	case entity.Open, entity.InProgressTicket, entity.Closed:
		return true
	default:
		return false
	}
}

func (s *Service) DeleteTicket(ctx *gin.Context, id string) error {
	ticket, err := s.Repo.GetTicket(ctx, id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteTicket(ctx, id, ticket); err != nil {
		return err
	}

	return nil
}
func (s *Service) SearchTicket(ctx *gin.Context, query string) (*[]entity.Ticket, error) {
	tickets, err := s.Repo.SearchTicket(ctx, query)
	if err != nil {
		return tickets, err
	}

	return tickets, nil
}
