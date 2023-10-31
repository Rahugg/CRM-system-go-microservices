package service

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetTickets(ctx *gin.Context) (*[]entity.Ticket, error) {
	tickets, err := s.Repo.GetTickets(ctx)

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

	ticket.IssueDescription = newTicket.IssueDescription
	ticket.Status = newTicket.Status
	ticket.ContactID = newTicket.ContactID
	ticket.AssignedTo = newTicket.AssignedTo

	if err = s.Repo.SaveTicket(ctx, ticket); err != nil {
		return err
	}

	return nil
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
