package repository

import (
	"crm_system/internal/auth/entity"
	"errors"
	"fmt"
	"strings"
)

func (r *AuthRepo) GetUserByIdWithPreload(id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("the user belonging to this token no logger exists")
	}
	return &user, nil
}
func (r *AuthRepo) GetUser(id string) (*entity.User, error) {
	var user *entity.User
	if err := r.DB.Where("id = ?", fmt.Sprint(id)).First(&user).Error; err != nil {
		return nil, errors.New("the user belonging to this token no logger exists")
	}
	return user, nil
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

func (r *AuthRepo) SaveUser(user *entity.User) error {
	if err := r.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *AuthRepo) GetAllUsers() (*[]entity.User, error) {
	var users *[]entity.User

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *AuthRepo) DeleteUser(id string, user *entity.User) error {
	if err := r.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthRepo) SearchUser(query string) (*[]entity.User, error) {
	var users *[]entity.User

	if err := r.DB.Where("first_name ILIKE ?", "%"+query+"%").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
func (r *AuthRepo) CreateUserCode(id string, code string) error {
	userCode := entity.UserCode{
		UserID: id,
		Code:   code,
	}
	err := r.DB.Create(&userCode).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepo) ConfirmUser(code string) error {
	var userCode *entity.UserCode

	if err := r.DB.Where("code = ?", fmt.Sprint(code)).First(&userCode).Error; err != nil {
		return err
	}

	userID := userCode.UserID

	if err := r.DB.Where("code = ?", userCode.Code).Delete(userCode).Error; err != nil {
		return err
	}

	user, err := r.GetUser(userID)
	if err != nil {
		return err
	}
	user.IsConfirmed = true

	err = r.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepo) SortUsers(users *[]entity.User, sortBy, sortOrder string) (*[]entity.User, error) {
	if err := r.DB.Order(sortBy + " " + sortOrder).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *AuthRepo) FilterUsersByAge(users *[]entity.User, age string) (*[]entity.User, error) {
	if err := r.DB.Where("age > ?", age).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
