package service

import (
	"crm_system/internal/auth/entity"
	"github.com/gin-gonic/gin"
)

type (
	AuthService interface {
		SignUp(ctx *gin.Context, payload *entity.SignUpInput, roleId uint, verified bool, provider string) (interface{}, error)
		SignIn(payload *entity.SignInInput) (*entity.SignInResult, error)
	}
	UserService interface {
		GetUsers(ctx *gin.Context) (*[]entity.User, error)
		GetUser(ctx *gin.Context, id string) (*entity.User, error)
		CreateUser(ctx *gin.Context, user entity.User) error
		UpdateUser(ctx *gin.Context, newUser entity.User, id string) error
		DeleteUser(ctx *gin.Context, id string) error
	}
)
