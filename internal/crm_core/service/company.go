package service

import (
	"crm_system/internal/crm_core/entity"
)

func (s *Service) GetCompanies(sortBy, sortOrder, phone string) (*[]entity.Company, error) {
	companies, err := s.Repo.GetCompanies()
	if err != nil {
		return nil, err
	}
	if phone != "" {
		companies, err = s.filterCompaniesByPhone(companies, phone)
		if err != nil {
			return nil, err
		}
	}

	if sortBy != "" {
		companies, err = s.sortCompanies(companies, sortBy, sortOrder)
		if err != nil {
			return nil, err
		}
	}

	return companies, nil
}

func (s *Service) sortCompanies(companies *[]entity.Company, sortBy, sortOrder string) (*[]entity.Company, error) {
	companies, err := s.Repo.SortCompanies(companies, sortBy, sortOrder)

	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *Service) filterCompaniesByPhone(companies *[]entity.Company, phone string) (*[]entity.Company, error) {
	companies, err := s.Repo.FilterCompaniesByPhone(companies, phone)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (s *Service) GetCompany(id string) (*entity.Company, error) {
	company, err := s.Repo.GetCompany(id)

	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *Service) CreateCompany(company entity.Company) error {
	if err := s.Repo.CreateCompany(&company); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateCompany(newCompany entity.NewCompany, id string) error {
	company, err := s.Repo.GetCompany(id)
	if err != nil {
		return err
	}

	if newCompany.Name != "" {
		company.Name = newCompany.Name
	}

	if newCompany.Address != "" {
		company.Address = newCompany.Address
	}

	if newCompany.Phone != "" {
		company.Phone = newCompany.Phone
	}

	if err = s.Repo.SaveCompany(company); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteCompany(id string) error {
	company, err := s.Repo.GetCompany(id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteCompany(id, company); err != nil {
		return err
	}

	return nil
}

func (s *Service) SearchCompany(query string) (*[]entity.Company, error) {
	companies, err := s.Repo.SearchCompany(query)
	if err != nil {
		return companies, err
	}

	return companies, nil
}
