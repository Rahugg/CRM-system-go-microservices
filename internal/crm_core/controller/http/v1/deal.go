package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/service"
	"crm_system/pkg/crm_core/logger"
	"github.com/gin-gonic/gin"
)

type dealRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newDealRoutes(handler *gin.RouterGroup, s *service.Service, l *logger.Logger, MW *middleware.Middleware) {
	r := &dealRoutes{s, l}

	dealHandler := handler.Group("/deal")
	{
		//middleware for users
		dealHandler.GET("/", r.getDeals)
		dealHandler.GET("/:id", r.getDeal)
		dealHandler.POST("/", r.createDeal)
		dealHandler.PUT("/:id", r.updateDeal)
		dealHandler.DELETE("/:id", r.deleteDeal)
	}
}

func (dr *dealRoutes) getDeals(ctx *gin.Context) {

}
func (dr *dealRoutes) getDeal(ctx *gin.Context) {

}
func (dr *dealRoutes) createDeal(ctx *gin.Context) {

}
func (dr *dealRoutes) updateDeal(ctx *gin.Context) {

}
func (dr *dealRoutes) deleteDeal(ctx *gin.Context) {

}
