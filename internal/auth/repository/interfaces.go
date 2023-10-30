package repository

import "crm_system/internal/auth/entity"

type AuthRepository interface {
	GetUserByIdWithPreload(id string) (*entity.User, error)
	GetUserById(id string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUsersByRole(roleId uint) (*[]entity.User, error)
	CreateUser(user *entity.User) error
	GetRoleById(id uint) (*entity.Role, error)
}
