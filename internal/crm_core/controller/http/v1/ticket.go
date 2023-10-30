package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/service"
	"crm_system/pkg/crm_core/logger"
	"github.com/gin-gonic/gin"
)

type ticketRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newTicketRoutes(handler *gin.RouterGroup, s *service.Service, l *logger.Logger, MW *middleware.Middleware) {
	r := &ticketRoutes{s, l}

	ticketHandler := handler.Group("/ticket")
	{
		//middleware for users
		ticketHandler.GET("/", r.getTickets)
		ticketHandler.GET("/:id", r.getTicket)
		ticketHandler.POST("/", r.createTicket)
		ticketHandler.PUT("/:id", r.updateTicket)
		ticketHandler.DELETE("/:id", r.deleteTicket)
	}
}

func (tr *ticketRoutes) getTickets(ctx *gin.Context) {

}
func (tr *ticketRoutes) getTicket(ctx *gin.Context) {

}
func (tr *ticketRoutes) createTicket(ctx *gin.Context) {

}
func (tr *ticketRoutes) updateTicket(ctx *gin.Context) {

}
func (tr *ticketRoutes) deleteTicket(ctx *gin.Context) {

}
