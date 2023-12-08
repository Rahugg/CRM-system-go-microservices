package service

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetContacts(ctx *gin.Context, sortBy, sortOrder, phone string) (*[]entity.Contact, error) {
	contacts, err := s.Repo.GetContacts(ctx)

	if phone != "" {
		contacts, err = s.filterContactsByPhone(contacts, phone)
		if err != nil {
			return nil, err
		}
	}

	if sortBy != "" {
		contacts, err = s.sortContacts(contacts, sortBy, sortOrder)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}
	return contacts, nil
}
func (s *Service) sortContacts(contacts *[]entity.Contact, sortBy, sortOrder string) (*[]entity.Contact, error) {
	contacts, err := s.Repo.SortContacts(contacts, sortBy, sortOrder)

	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s *Service) filterContactsByPhone(contacts *[]entity.Contact, phone string) (*[]entity.Contact, error) {
	contacts, err := s.Repo.FilterContactsByPhone(contacts, phone)
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

	if newContact.FirstName != "" {
		contact.FirstName = newContact.FirstName
	}

	if newContact.LastName != "" {
		contact.LastName = newContact.LastName
	}

	if newContact.Phone != "" {
		contact.Phone = newContact.Phone
	}

	if newContact.Email != "" {
		contact.Email = newContact.Email
	}

	if newContact.CompanyID != 0 {
		contact.CompanyID = newContact.CompanyID
	} else {
		contact.CompanyID = 0
	}

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

func (s *Service) SearchContact(ctx *gin.Context, query string) (*[]entity.Contact, error) {
	contacts, err := s.Repo.SearchContact(ctx, query)
	if err != nil {
		return contacts, err
	}

	return contacts, nil
}
