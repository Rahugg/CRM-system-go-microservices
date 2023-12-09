package service

import (
	"crm_system/internal/auth/entity"
	"github.com/gin-gonic/gin"
)

type (
	AuthService interface {
		SignUp(payload *entity.SignUpInput, roleId uint, provider string) (string, error)
		SignIn(payload *entity.SignInInput) (*entity.SignInResult, error)
	}
	UserService interface {
		GetUsers(sortBy, sortOrder, age string) (*[]entity.User, error)
		GetUser(id string) (*entity.User, error)
		CreateUser(user *entity.User) error
		UpdateUser(newUser *entity.User, id string) error
		DeleteUser(id string) error
		GetMe(ctx *gin.Context) (interface{}, error)
		UpdateMe(ctx *gin.Context, newUser *entity.User) error
		SearchUser(query string) (*[]entity.User, error)
		CreateUserCode(id string) error
		ConfirmUser(code string) error
	}
)
