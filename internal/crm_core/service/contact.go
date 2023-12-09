package service

import (
	"crm_system/internal/crm_core/entity"
)

func (s *Service) GetContacts(sortBy, sortOrder, phone string) (*[]entity.Contact, error) {
	contacts, err := s.Repo.GetContacts()

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

func (s *Service) GetContact(id string) (*entity.Contact, error) {
	contact, err := s.Repo.GetContact(id)

	if err != nil {
		return nil, err
	}

	return contact, nil
}
func (s *Service) CreateContact(contact entity.Contact) error {
	if err := s.Repo.CreateContact(&contact); err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateContact(newContact entity.Contact, id string) error {
	contact, err := s.Repo.GetContact(id)
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

	if err = s.Repo.SaveContact(contact); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteContact(id string) error {
	contact, err := s.Repo.GetContact(id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteContact(id, contact); err != nil {
		return err
	}

	return nil
}

func (s *Service) SearchContact(query string) (*[]entity.Contact, error) {
	contacts, err := s.Repo.SearchContact(query)
	if err != nil {
		return contacts, err
	}

	return contacts, nil
}
