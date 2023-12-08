package middleware

import (
	"crm_system/config/crm_core"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/metrics"
	"crm_system/internal/crm_core/repository"
	"crm_system/internal/crm_core/transport"
	pb "crm_system/pkg/auth/authservice/gw"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func getUser(resp *pb.ResponseJSON) *entity.User {
	userBuilder := entity.NewUser()
	user := userBuilder.
		SetFirstName(resp.User.FirstName).
		SetLastName(resp.User.LastName).
		SetAge(uint64(resp.User.Age)).
		SetPhone(resp.User.Phone).
		SetRoleID(uint(resp.User.RoleID)).
		SetEmail(resp.User.Email).
		SetProvider(resp.User.Provider).
		SetPassword(resp.User.Password).
		Build()
	return &user
}

func (m *Middleware) DeserializeUser(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := validBearer(ctx)

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &entity.CustomResponse{
				Status:  -1,
				Message: "You are not logged in",
			})
			return
		}

		resp, err := m.ValidateGrpcTransport.ValidateTransport(ctx.Request.Context(), accessToken, roles...)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &entity.CustomResponse{
				Status:  -1,
				Message: err.Error(),
			})
			return
		}

		user := getUser(resp)

		ctx.Set("currentUser", user)
		ctx.Set("currentRole", resp.Role)
		ctx.Next()
	}
}

func (m *Middleware) MetricsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		path := ctx.Request.URL.Path

		statusString := strconv.Itoa(ctx.Writer.Status())

		metrics.HttpResponseTime.WithLabelValues(path, statusString, ctx.Request.Method).Observe(time.Since(start).Seconds())
		metrics.HttpRequestsTotalCollector.WithLabelValues(path, statusString, ctx.Request.Method).Inc()
	}
}
