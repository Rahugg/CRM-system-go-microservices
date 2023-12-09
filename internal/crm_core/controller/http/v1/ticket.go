package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ticketRoutes struct {
	s *service.Service
}

func newTicketRoutes(handler *gin.RouterGroup, s *service.Service, MW *middleware.Middleware) {
	r := &ticketRoutes{s: s}

	ticketHandler := handler.Group("/ticket")
	{
		//middleware for users
		ticketHandler.Use(MW.MetricsHandler())
		ticketHandler.Use(MW.DeserializeUser("admin", "manager"))

		ticketHandler.GET("/", r.getTickets)
		ticketHandler.GET("/:id", r.getTicket)
		ticketHandler.POST("/", r.createTicket)
		ticketHandler.PUT("/:id", r.updateTicket)
		ticketHandler.DELETE("/:id", r.deleteTicket)
		ticketHandler.GET("/search", r.searchTicket)
	}
}

// getTickets godoc
// @Summary Получить список билетов
// @Description Получить список билетов
// @Tags ticket
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param sortBy query string false "sortBy"
// @Param sortOrder query string false "sortOrder"
// @Param status query string false "filter by status"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 404 {object} entity.CustomResponse
// @Router /v1/ticket/ [get]
func (tr *ticketRoutes) getTickets(ctx *gin.Context) {
	sortBy := ctx.Query("sortBy")
	sortOrder := ctx.Query("sortOrder")
	status := ctx.Query("status")
	tickets, err := tr.s.GetTickets(sortBy, sortOrder, status)

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

// getTicket godoc
// @Summary Получить билет по id
// @Description Получить билет по id
// @Tags ticket
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id ticket"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/ticket/{id} [get]
func (tr *ticketRoutes) getTicket(ctx *gin.Context) {
	id := ctx.Param("id")

	ticket, err := tr.s.GetTicket(id)

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

// createTicket godoc
// @Summary Создать Билет
// @Description Создать Билет
// @Tags ticket
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ticketInput body entity.Ticket true "Create ticket"
// @Success 201 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/ticket/ [post]
func (tr *ticketRoutes) createTicket(ctx *gin.Context) {
	var ticket entity.Ticket

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := tr.s.CreateTicket(ticket); err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}

// updateTicket godoc
// @Summary Редактировать билет по id
// @Description Редактировать билет по id
// @Tags ticket
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id ticket"
// @Param inputTicket body entity.Ticket true "Update Ticket"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/ticket/{id} [put]
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

	if err := tr.s.UpdateTicket(newTicket, id); err != nil {
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

// deleteTicket godoc
// @Summary Удалить билет по id
// @Description Удалить билет по id
// @Tags ticket
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id ticket"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/ticket/{id} [delete]
func (tr *ticketRoutes) deleteTicket(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := tr.s.DeleteTicket(id); err != nil {
		ctx.JSON(http.StatusNoContent, &entity.CustomResponse{
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

// searchTicket godoc
// @Summary Поиск билета по имени
// @Description Поиск билета по имени
// @Tags ticket
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param searchQuery query string true "query"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/ticket/search [get]
func (tr *ticketRoutes) searchTicket(ctx *gin.Context) {
	query := ctx.Query("searchQuery")
	tickets, err := tr.s.SearchTicket(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponseWithData{
			Status:  -1,
			Message: "Not found",
			Data:    tickets,
		})
		return
	}
	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    tickets,
	})
}
