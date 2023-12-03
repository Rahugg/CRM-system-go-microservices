package repository

import (
	"crm_system/internal/crm_core/entity"
	"github.com/gin-gonic/gin"
)

func (r *CRMSystemRepo) GetDeals(ctx *gin.Context) (*[]entity.Deal, error) {
	var deals *[]entity.Deal

	if err := r.DB.Find(&deals).Error; err != nil {
		return nil, err
	}

	return deals, nil
}

func (r *CRMSystemRepo) GetDeal(ctx *gin.Context, id string) (*entity.Deal, error) {
	var deal *entity.Deal

	if err := r.DB.First(&deal, id).Error; err != nil {
		return nil, err
	}

	return deal, nil
}

func (r *CRMSystemRepo) CreateDeal(ctx *gin.Context, deal *entity.Deal) error {
	if err := r.DB.Create(&deal).Error; err != nil {
		return err
	}

	return nil
}

func (r *CRMSystemRepo) SaveDeal(ctx *gin.Context, newDeal *entity.Deal) error {
	if err := r.DB.Save(&newDeal).Error; err != nil {
		return err
	}

	return nil
}

func (r *CRMSystemRepo) DeleteDeal(ctx *gin.Context, id string, deal *entity.Deal) error {
	if err := r.DB.Where("id = ?", id).Delete(deal).Error; err != nil {
		return err
	}
	return nil
}

func (r *CRMSystemRepo) SearchDeal(ctx *gin.Context, query string) (*[]entity.Deal, error) {
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
