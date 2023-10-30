package service

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetContacts(ctx *gin.Context) (*[]entity.Contact, error) {
	contacts, err := s.Repo.GetContacts(ctx)

	if err != nil {
		return nil, err
	}
	return contacts, nil
}
func (s *Service) GetContact(ctx *gin.Context, id string) (*entity.Contact, error) {
	contact, err := s.Repo.GetContact(ctx, id)

	if err != nil {
		return nil, err
	}

	return contact, nil
}
func (s *Service) CreateContact(ctx *gin.Context, contact entity.Contact) error {
	if err := s.Repo.CreateContact(ctx, &contact); err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateContact(ctx *gin.Context, newContact entity.Contact, id string) error {
	contact, err := s.Repo.GetContact(ctx, id)
	if err != nil {
		return err
	}

	contact.FirstName = newContact.FirstName
	contact.LastName = newContact.LastName
	contact.Phone = newContact.Phone
	contact.Email = newContact.Email
	contact.CompanyID = newContact.CompanyID

	if err = s.Repo.SaveContact(ctx, contact); err != nil {
		return err
	}

	return nil
}
func (s *Service) DeleteContact(ctx *gin.Context, id string) error {
	contact, err := s.Repo.GetContact(ctx, id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteContact(ctx, id, contact); err != nil {
		return err
	}

	return nil
}
