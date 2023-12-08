package repository

import (
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/metrics"
	"github.com/gin-gonic/gin"
)

func (r *CRMSystemRepo) GetContacts(ctx *gin.Context) (*[]entity.Contact, error) {
	var contacts *[]entity.Contact
	ok, fail := metrics.DatabaseQueryTime("GetContacts")
	defer fail()

	if err := r.DB.Find(&contacts).Error; err != nil {
		return nil, err
	}

	ok()

	return contacts, nil
}
func (r *CRMSystemRepo) GetContact(ctx *gin.Context, id string) (*entity.Contact, error) {
	var contact *entity.Contact

	if err := r.DB.First(&contact, id).Error; err != nil {
		return nil, err
	}

	return contact, nil
}
func (r *CRMSystemRepo) CreateContact(ctx *gin.Context, contact *entity.Contact) error {
	if err := r.DB.Create(&contact).Error; err != nil {
		return err
	}

	return nil
}
func (r *CRMSystemRepo) SaveContact(ctx *gin.Context, newContact *entity.Contact) error {
	if err := r.DB.Save(&newContact).Error; err != nil {
		return err
	}

	return nil
}
func (r *CRMSystemRepo) DeleteContact(ctx *gin.Context, id string, contact *entity.Contact) error {
	if err := r.DB.Where("id = ?", id).Delete(contact).Error; err != nil {
		return err
	}
	return nil
}

func (r *CRMSystemRepo) SearchContact(ctx *gin.Context, query string) (*[]entity.Contact, error) {
	var contacts *[]entity.Contact

	if err := r.DB.Where("first_name ILIKE ?", "%"+query+"%").Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *CRMSystemRepo) SortContacts(contacts *[]entity.Contact, sortBy, sortOrder string) (*[]entity.Contact, error) {
	if err := r.DB.Order(sortBy + " " + sortOrder).Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *CRMSystemRepo) FilterContactsByPhone(contacts *[]entity.Contact, phone string) (*[]entity.Contact, error) {
	if err := r.DB.Where("phone = ?", phone).Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}
