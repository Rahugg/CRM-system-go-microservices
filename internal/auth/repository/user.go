package repository

import (
	"errors"
	"fmt"
	"crm_system/internal/auth/entity"
	"strings"
)

func (r *AuthRepo) GetUserByIdWithPreload(id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("the user belonging to this token no logger exists")
	}
	return &user, nil
}
func (r *AuthRepo) GetUserById(id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", fmt.Sprint(id)).First(&user).Error; err != nil {
		return nil, errors.New("the user belonging to this token no logger exists")
	}
	return &user, nil
}

func (r *AuthRepo) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Preload("Role").Where("email = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepo) GetUsersByRole(roleId uint) (*[]entity.User, error) {
	var users []entity.User
	if err := r.DB.Preload("Role").Where("role_id = ? ", roleId).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *AuthRepo) CreateUser(user *entity.User) error {
	err := r.DB.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique") {
			return errors.New("user with that email already exists")
		}
		return err
	}
	return nil
}

func (r *AuthRepo) SaveUser(user *entity.User) {
	r.DB.Save(user)
}
