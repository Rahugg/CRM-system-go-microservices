package service

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetDeals(ctx *gin.Context) (*[]entity.Deal, error) {
	deals, err := s.Repo.GetDeals(ctx)

	if err != nil {
		return nil, err
	}
	return deals, nil
}

func (s *Service) GetDeal(ctx *gin.Context, id string) (*entity.Deal, error) {
	deal, err := s.Repo.GetDeal(ctx, id)

	if err != nil {
		return nil, err
	}

	return deal, nil
}
func (s *Service) CreateDeal(ctx *gin.Context, deal entity.Deal) error {
	if err := s.Repo.CreateDeal(ctx, &deal); err != nil {
		return err
	}

	return nil
}
func (s *Service) UpdateDeal(ctx *gin.Context, newDeal entity.Deal, id string) error {
	deal, err := s.Repo.GetDeal(ctx, id)
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

	if err = s.Repo.SaveDeal(ctx, deal); err != nil {
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

func (s *Service) DeleteDeal(ctx *gin.Context, id string) error {
	deal, err := s.Repo.GetDeal(ctx, id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteDeal(ctx, id, deal); err != nil {
		return err
	}

	return nil
}