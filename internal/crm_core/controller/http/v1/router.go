package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/service"
	"crm_system/pkg/crm_core/cache"
	"crm_system/pkg/crm_core/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(handler *gin.Engine, s *service.Service, l *logger.Logger, MW *middleware.Middleware, cc cache.Contact) {

	// Health Check
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello to auth service",
		})
	})

	h := handler.Group("/v1")
	{
		newCompanyRoutes(h, s, l, MW)
		newContactRoutes(h, s, l, MW, cc)
		newDealRoutes(h, s, l, MW)
		newTaskRoutes(h, s, l, MW)
		newTicketRoutes(h, s, l, MW)
	}
}

//customer relation
//open source crm-system
