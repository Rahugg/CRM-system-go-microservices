package repository

import (
	"crm_system/internal/auth/entity"
	"github.com/gin-gonic/gin"
)

func (r *AuthRepo) GetRoleById(ctx *gin.Context, id uint) (*entity.Role, error) {
	var role entity.Role
	if err := r.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
