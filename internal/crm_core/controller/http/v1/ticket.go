package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/service"
	"crm_system/pkg/crm_core/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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
	tickets, err := tr.s.GetTickets(ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    tickets,
	})
}
func (tr *ticketRoutes) getTicket(ctx *gin.Context) {
	id := ctx.Param("id")

	ticket, err := tr.s.GetTicket(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    ticket,
	})
}
func (tr *ticketRoutes) createTicket(ctx *gin.Context) {
	var ticket entity.Ticket

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := tr.s.CreateTicket(ctx, ticket); err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}
func (tr *ticketRoutes) updateTicket(ctx *gin.Context) {
	id := ctx.Param("id")

	var newTicket entity.Ticket

	if err := ctx.ShouldBindJSON(&newTicket); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := tr.s.UpdateTicket(ctx, newTicket, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}
func (tr *ticketRoutes) deleteTicket(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := tr.s.DeleteTicket(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}
