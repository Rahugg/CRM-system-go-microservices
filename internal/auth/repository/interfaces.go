package repository

import (
	"crm_system/internal/auth/entity"
	"github.com/gin-gonic/gin"
)

type (
	AuthRepository interface {
		GetUserByIdWithPreload(ctx *gin.Context, id string) (*entity.User, error)
		GetUser(ctx *gin.Context, id string) (*entity.User, error)
		GetUserByEmail(ctx *gin.Context, email string) (*entity.User, error)
		GetUsersByRole(ctx *gin.Context, roleId uint) (*[]entity.User, error)
		CreateUser(ctx *gin.Context, user *entity.User) error
		GetRoleById(ctx *gin.Context, id uint) (*entity.Role, error)
		SaveUser(ctx *gin.Context, user *entity.User) error
	}
	UserRepository interface {
		GetUsers(ctx *gin.Context) (*[]entity.User, error)
		DeleteDeal(ctx *gin.Context, id string, deal *entity.User) error
	}
)
