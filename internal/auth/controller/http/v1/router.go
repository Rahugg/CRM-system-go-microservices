package v1

import (
	"crm_system/internal/auth/controller/http/middleware"
	"crm_system/internal/auth/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(handler *gin.Engine, s *service.Service, MW *middleware.Middleware) {

	// Health Check
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello to auth service",
		})
	})

	h := handler.Group("/v1")
	{
		newUserRoutes(h, s, MW)
	}
}
