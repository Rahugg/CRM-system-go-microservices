package middleware

import (
	"crm_system/config/auth"
	"crm_system/internal/auth/entity"
	"crm_system/internal/auth/repository"
	"crm_system/pkg/auth/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type Middleware struct {
	Repo   *repository.AuthRepo
	Config *auth.Configuration
}

func New(repo *repository.AuthRepo, config *auth.Configuration) *Middleware {
	return &Middleware{
		Repo:   repo,
		Config: config,
	}
}

func (m *Middleware) CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before")

		c.Next()

		log.Println("after")
	}
}

func validBearer(ctx *gin.Context) string {
	var accessToken string

	cookie, err := ctx.Cookie("access_token")
	authorizationHeader := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else if err == nil {
		accessToken = cookie
	}

	return accessToken
}

func (m *Middleware) DeserializeUser(roles ...interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := validBearer(ctx)

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &entity.CustomResponse{
				Status:  -1,
				Message: "You are not logged in",
			})
			return
		}

		sub, err := utils.ValidateToken(accessToken, m.Config.Jwt.AccessPrivateKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &entity.CustomResponse{
				Status:  -2,
				Message: err.Error(),
			})
			return
		}

		user, err := m.Repo.GetUserByIdWithPreload(fmt.Sprint(sub))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, &entity.CustomResponse{
				Status:  -3,
				Message: "the user belonging to this token no logger exists",
			})
			return
		}

		role, _ := m.Repo.GetRoleById(user.RoleID)

		for _, Role := range roles {
			if role.Name == Role || Role == "any" {
				ctx.Set("currentUser", user)
				ctx.Set("currentRole", role)
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, &entity.CustomResponse{
			Status:  -4,
			Message: "User must have roles: " + fmt.Sprintf("%v", roles),
		})
	}
}
