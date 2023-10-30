package repository

import (
	"crm_system/internal/auth/entity"
)

func (r *AuthRepo) GetRoleById(id uint) (*entity.Role, error) {
	var role entity.Role
	if err := r.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
