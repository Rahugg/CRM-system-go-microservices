package service

import (
	"crm_system/internal/crm_core/entity"
)

func (s *Service) GetDeals(sortBy, sortOrder, status string) (*[]entity.Deal, error) {
	deals, err := s.Repo.GetDeals()
	if err != nil {
		return nil, err
	}

	if status != "" {
		deals, err = s.filterDealsByStatus(deals, status)
		if err != nil {
			return nil, err
		}
	}

	if sortBy != "" {
		deals, err = s.sortDeals(deals, sortBy, sortOrder)
		if err != nil {
			return nil, err
		}
	}

	return deals, nil
}
func (s *Service) sortDeals(deals *[]entity.Deal, sortBy, sortOrder string) (*[]entity.Deal, error) {
	deals, err := s.Repo.SortDeals(deals, sortBy, sortOrder)

	if err != nil {
		return nil, err
	}

	return deals, nil
}

func (s *Service) filterDealsByStatus(deals *[]entity.Deal, status string) (*[]entity.Deal, error) {
	deals, err := s.Repo.FilterDealsByStatus(deals, status)
	if err != nil {
		return nil, err
	}
	return deals, nil
}

func (s *Service) GetDeal(id string) (*entity.Deal, error) {
	deal, err := s.Repo.GetDeal(id)

	if err != nil {
		return nil, err
	}

	return deal, nil
}
func (s *Service) CreateDeal(deal entity.Deal) error {
	if err := s.Repo.CreateDeal(&deal); err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateDeal(newDeal entity.Deal, id string) error {
	deal, err := s.Repo.GetDeal(id)
	if err != nil {
		return err
	}

	if newDeal.Title != "" {
		deal.Title = newDeal.Title
	}

	if newDeal.Value != 0 {
		deal.Value = newDeal.Value
	}

	if isValidDealStatus(newDeal.Status) {
		deal.Status = newDeal.Status
	}

	if newDeal.ContactID != 0 {
		deal.ContactID = newDeal.ContactID
	}

	if newDeal.RepID != 0 {
		deal.RepID = newDeal.RepID
	}

	if err = s.Repo.SaveDeal(deal); err != nil {
		return err
	}

	return nil
}
func isValidDealStatus(status entity.StatusDeal) bool {
	switch status {
	case entity.Initiated, entity.InProgress, entity.ClosedWon, entity.ClosedLost:
		return true
	default:
		return false
	}
}

func (s *Service) DeleteDeal(id string) error {
	deal, err := s.Repo.GetDeal(id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteDeal(id, deal); err != nil {
		return err
	}

	return nil
}

func (s *Service) SearchDeal(query string) (*[]entity.Deal, error) {
	deals, err := s.Repo.SearchDeal(query)
	if err != nil {
		return deals, err
	}

	return deals, nil
}
