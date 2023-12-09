package repository

import (
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/metrics"
)

func (r *CRMSystemRepo) GetDeals() (*[]entity.Deal, error) {
	var deals *[]entity.Deal
	ok, fail := metrics.DatabaseQueryTime("GetDeals")
	defer fail()

	if err := r.DB.Find(&deals).Error; err != nil {
		return nil, err
	}

	ok()
	return deals, nil
}

func (r *CRMSystemRepo) GetDeal(id string) (*entity.Deal, error) {
	var deal *entity.Deal

	if err := r.DB.First(&deal, id).Error; err != nil {
		return nil, err
	}

	return deal, nil
}

func (r *CRMSystemRepo) CreateDeal(deal *entity.Deal) error {
	if err := r.DB.Create(&deal).Error; err != nil {
		return err
	}

	return nil
}

func (r *CRMSystemRepo) SaveDeal(newDeal *entity.Deal) error {
	if err := r.DB.Save(&newDeal).Error; err != nil {
		return err
	}

	return nil
}

func (r *CRMSystemRepo) DeleteDeal(id string, deal *entity.Deal) error {
	if err := r.DB.Where("id = ?", id).Delete(deal).Error; err != nil {
		return err
	}
	return nil
}

func (r *CRMSystemRepo) SearchDeal(query string) (*[]entity.Deal, error) {
	var deals *[]entity.Deal

	if err := r.DB.Where("title ILIKE ?", "%"+query+"%").Find(&deals).Error; err != nil {
		return nil, err
	}

	return deals, nil
}
func (r *CRMSystemRepo) SortDeals(deals *[]entity.Deal, sortBy, sortOrder string) (*[]entity.Deal, error) {
	if err := r.DB.Order(sortBy + " " + sortOrder).Find(&deals).Error; err != nil {
		return nil, err
	}

	return deals, nil
}

func (r *CRMSystemRepo) FilterDealsByStatus(deals *[]entity.Deal, status string) (*[]entity.Deal, error) {
	if err := r.DB.Where("status = ?", status).Find(&deals).Error; err != nil {
		return nil, err
	}

	return deals, nil
}
