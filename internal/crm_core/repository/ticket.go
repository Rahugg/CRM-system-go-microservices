package repository

import (
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/metrics"
)

func (r *CRMSystemRepo) GetTickets() (*[]entity.Ticket, error) {
	var tickets *[]entity.Ticket

	ok, fail := metrics.DatabaseQueryTime("GetTickets")
	defer fail()

	if err := r.DB.Find(&tickets).Error; err != nil {
		return nil, err
	}

	ok()

	return tickets, nil
}
func (r *CRMSystemRepo) GetTicket(id string) (*entity.Ticket, error) {
	var ticket *entity.Ticket

	if err := r.DB.First(&ticket, id).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}
func (r *CRMSystemRepo) CreateTicket(ticket *entity.Ticket) error {
	if err := r.DB.Create(&ticket).Error; err != nil {
		return err
	}

	return nil
}
func (r *CRMSystemRepo) SaveTicket(newTicket *entity.Ticket) error {
	if err := r.DB.Save(&newTicket).Error; err != nil {
		return err
	}

	return nil
}
func (r *CRMSystemRepo) DeleteTicket(id string, ticket *entity.Ticket) error {
	if err := r.DB.Where("id = ?", id).Delete(ticket).Error; err != nil {
		return err
	}
	return nil
}

func (r *CRMSystemRepo) SearchTicket(query string) (*[]entity.Ticket, error) {
	var tickets *[]entity.Ticket

	if err := r.DB.Where("issue_description ILIKE ?", "%"+query+"%").Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}
func (r *CRMSystemRepo) SortTickets(tickets *[]entity.Ticket, sortBy, sortOrder string) (*[]entity.Ticket, error) {
	if err := r.DB.Order(sortBy + " " + sortOrder).Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *CRMSystemRepo) FilterTicketsByStatus(tickets *[]entity.Ticket, status string) (*[]entity.Ticket, error) {
	if err := r.DB.Where("status = ?", status).Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}
