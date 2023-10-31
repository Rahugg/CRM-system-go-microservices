package repository

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

func (r *CRMSystemRepo) GetTickets(ctx *gin.Context) (*[]entity.Ticket, error) {
	var tickets *[]entity.Ticket

	if err := r.DB.Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}
func (r *CRMSystemRepo) GetTicket(ctx *gin.Context, id string) (*entity.Ticket, error) {
	var ticket *entity.Ticket

	if err := r.DB.First(&ticket, id).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}
func (r *CRMSystemRepo) CreateTicket(ctx *gin.Context, ticket *entity.Ticket) error {
	if err := r.DB.Create(&ticket).Error; err != nil {
		return err
	}

	return nil
}
func (r *CRMSystemRepo) SaveTicket(ctx *gin.Context, newTicket *entity.Ticket) error {
	if err := r.DB.Save(&newTicket).Error; err != nil {
		return err
	}

	return nil
}
func (r *CRMSystemRepo) DeleteTicket(ctx *gin.Context, id string, ticket *entity.Ticket) error {
	if err := r.DB.Where("id = ?", id).Delete(ticket).Error; err != nil {
		return err
	}
	return nil
}
