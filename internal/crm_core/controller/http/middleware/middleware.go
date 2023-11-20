package middleware

import (
	"crm_system/config/crm_core"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/repository"
	"crm_system/internal/crm_core/transport"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type Middleware struct {
	Repo                  *repository.CRMSystemRepo
	Config                *crm_core.Configuration
	ValidateGrpcTransport *transport.ValidateGrpcTransport
}

func New(repo *repository.CRMSystemRepo, config *crm_core.Configuration, validateGrpcTransport *transport.ValidateGrpcTransport) *Middleware {
	return &Middleware{
		Repo:                  repo,
		Config:                config,
		ValidateGrpcTransport: validateGrpcTransport,
	}
}

func (m *Middleware) CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before")

		c.Next()

		log.Println("after")
	}
}

func (m *Middleware) DeserializeUser(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string

		cookie, err := ctx.Cookie("access_token")
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &entity.CustomResponse{
				Status:  -1,
				Message: "You are not logged in",
			})
			return
		}
		resp, err := m.ValidateGrpcTransport.Validate(ctx, accessToken, roles)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &entity.CustomResponse{
				Status:  -1,
				Message: err.Error(),
			})
			return
		}
		ctx.Set("currentUser", resp.User)
		ctx.Set("currentRole", resp.Role)
		ctx.Next()

	}
}
