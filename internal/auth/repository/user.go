package repository

import (
	"crm_system/internal/auth/entity"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func (r *AuthRepo) GetUserByIdWithPreload(ctx *gin.Context, id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("the user belonging to this token no logger exists")
	}
	return &user, nil
}
func (r *AuthRepo) GetUser(ctx *gin.Context, id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", fmt.Sprint(id)).First(&user).Error; err != nil {
		return nil, errors.New("the user belonging to this token no logger exists")
	}
	return &user, nil
}

func (r *AuthRepo) GetUserByEmail(ctx *gin.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Preload("Role").Where("email = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepo) GetUsersByRole(ctx *gin.Context, roleId uint) (*[]entity.User, error) {
	var users []entity.User
	if err := r.DB.Preload("Role").Where("role_id = ? ", roleId).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *AuthRepo) CreateUser(ctx *gin.Context, user *entity.User) error {
	err := r.DB.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique") {
			return errors.New("user with that email already exists")
		}
		return err
	}
	return nil
}

func (r *AuthRepo) SaveUser(ctx *gin.Context, user *entity.User) error {
	if err := r.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *AuthRepo) GetUsers(ctx *gin.Context) (*[]entity.User, error) {
	var users *[]entity.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
func (r *AuthRepo) DeleteUser(ctx *gin.Context, id string, user *entity.User) error {
	if err := r.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
