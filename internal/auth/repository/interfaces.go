package repository

import (
	"crm_system/internal/auth/entity"
)

type (
	AuthRepository interface {
		CreateUserCode(id string, code string) error
		ConfirmUser(code string) error
	}
	UserRepository interface {
		CreateUser(user *entity.User) error
		SaveUser(user *entity.User) error
		SortUsers(users *[]entity.User, sortBy, sortOrder string) (*[]entity.User, error)
		GetUserByIdWithPreload(id string) (*entity.User, error)
		GetUser(id string) (*entity.User, error)
		GetAllUsers() (*[]entity.User, error)
		GetUserByEmail(email string) (*entity.User, error)
		GetUsersByRole(roleId uint) (*[]entity.User, error)
		DeleteUser(id string, user *entity.User) error
		SearchUser(query string) (*[]entity.User, error)
		FilterUsersByAge(users *[]entity.User, age string) (*[]entity.User, error)
	}
	RoleRepository interface {
		GetRoleById(id uint) (*entity.Role, error)
	}
)
