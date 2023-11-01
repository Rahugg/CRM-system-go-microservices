package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/service"
	"crm_system/pkg/crm_core/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(handler *gin.Engine, s *service.Service, MW *middleware.Middleware, cc cache.Contact) {

	// Health Check
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello to auth service",
		})
	})

	h := handler.Group("/v1")
	{
		newCompanyRoutes(h, s, MW)
		newContactRoutes(h, s, MW, cc)
		newDealRoutes(h, s, MW)
		newTaskRoutes(h, s, MW)
		newTicketRoutes(h, s, MW)
	}
}

//customer relation
//open source crm-system
