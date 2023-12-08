package repository

import (
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/metrics"
	"github.com/gin-gonic/gin"
)

func (r *CRMSystemRepo) GetCompanies(ctx *gin.Context) (*[]entity.Company, error) {
	var companies *[]entity.Company
	ok, fail := metrics.DatabaseQueryTime("GetCompanies")
	defer fail()

	if err := r.DB.Find(&companies).Error; err != nil {
		return nil, err
	}

	ok()

	return companies, nil
}

func (r *CRMSystemRepo) GetCompany(ctx *gin.Context, id string) (*entity.Company, error) {
	var company *entity.Company

	if err := r.DB.First(&company, id).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (r *CRMSystemRepo) CreateCompany(ctx *gin.Context, company *entity.Company) error {
	if err := r.DB.Create(&company).Error; err != nil {
		return err
	}

	return nil
}

func (r *CRMSystemRepo) SaveCompany(ctx *gin.Context, newCompany *entity.Company) error {
	if err := r.DB.Save(&newCompany).Error; err != nil {
		return err
	}

	return nil
}

func (r *CRMSystemRepo) DeleteCompany(ctx *gin.Context, id string, company *entity.Company) error {
	if err := r.DB.Where("id = ?", id).Delete(company).Error; err != nil {
		return err
	}
	return nil
}
func (r *CRMSystemRepo) SearchCompany(ctx *gin.Context, query string) (*[]entity.Company, error) {
	var companies *[]entity.Company

	if err := r.DB.Where("name ILIKE ?", "%"+query+"%").Find(&companies).Error; err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *CRMSystemRepo) SortCompanies(companies *[]entity.Company, sortBy, sortOrder string) (*[]entity.Company, error) {
	if err := r.DB.Order(sortBy + " " + sortOrder).Find(&companies).Error; err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *CRMSystemRepo) FilterCompaniesByPhone(companies *[]entity.Company, phone string) (*[]entity.Company, error) {
	if err := r.DB.Where("phone = ?", phone).Find(&companies).Error; err != nil {
		return nil, err
	}

	return companies, nil
}
