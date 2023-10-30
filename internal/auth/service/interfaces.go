package service

import (
	"github.com/gin-gonic/gin"
	"crm_system/internal/auth/entity"
)

type (
	AuthService interface {
		SignUp(ctx *gin.Context, payload *entity.SignUpInput, roleId uint, verified bool, provider string) (interface{}, error)
		SignIn(payload *entity.SignInInput) (*entity.SignInResult, error)
	}
)
