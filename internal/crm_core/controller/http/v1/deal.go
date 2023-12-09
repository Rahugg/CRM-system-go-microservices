package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type dealRoutes struct {
	s *service.Service
}

func newDealRoutes(handler *gin.RouterGroup, s *service.Service, MW *middleware.Middleware) {
	r := &dealRoutes{s: s}

	dealHandler := handler.Group("/deal")
	{
		//middleware for users
		dealHandler.Use(MW.MetricsHandler())

		dealHandler.GET("/", MW.DeserializeUser("any"), r.getDeals)
		dealHandler.GET("/:id", MW.DeserializeUser("any"), r.getDeal)
		dealHandler.POST("/", MW.DeserializeUser("any"), r.createDeal)
		dealHandler.PUT("/:id", MW.DeserializeUser("any"), r.updateDeal)
		dealHandler.DELETE("/:id", MW.DeserializeUser("any"), r.deleteDeal)
		dealHandler.GET("/search", MW.DeserializeUser("any"), r.searchDeal)
	}
}

// getDeals godoc
// @Summary Получить список соглашений
// @Description Получить список соглашений
// @Tags deal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param sortBy query string false "sortBy"
// @Param sortOrder query string false "sortOrder"
// @Param status query string false "filter by status"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 404 {object} entity.CustomResponse
// @Router /v1/deal/ [get]
func (dr *dealRoutes) getDeals(ctx *gin.Context) {
	sortBy := ctx.Query("sortBy")
	sortOrder := ctx.Query("sortOrder")
	status := ctx.Query("status")
	deals, err := dr.s.GetDeals(sortBy, sortOrder, status)

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
		Data:    deals,
	})
}

// getDeal godoc
// @Summary Получить соглашение по id
// @Description Получить соглашение по id
// @Tags deal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id deal"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/deal/{id} [get]
func (dr *dealRoutes) getDeal(ctx *gin.Context) {
	id := ctx.Param("id")

	deal, err := dr.s.GetDeal(id)

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
		Data:    deal,
	})
}

// createDeal godoc
// @Summary Создать Соглашение
// @Description Создать Соглашение
// @Tags deal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param dealInput body entity.Deal true "Create deal"
// @Success 201 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/deal/ [post]
func (dr *dealRoutes) createDeal(ctx *gin.Context) {
	var deal entity.Deal

	if err := ctx.ShouldBindJSON(&deal); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := dr.s.CreateDeal(deal); err != nil {
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

// updateDeal godoc
// @Summary Редактировать соглашение по id
// @Description Редактировать соглашение по id
// @Tags deal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id deal"
// @Param newDeal body entity.Deal true "Update deal"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/deal/{id} [put]
func (dr *dealRoutes) updateDeal(ctx *gin.Context) {
	id := ctx.Param("id")

	var newDeal entity.Deal

	if err := ctx.ShouldBindJSON(&newDeal); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := dr.s.UpdateDeal(newDeal, id); err != nil {
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

// deleteDeal godoc
// @Summary Удалить соглашение по id
// @Description Удалить соглашение по id
// @Tags deal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id deal"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/deal/{id} [delete]
func (dr *dealRoutes) deleteDeal(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := dr.s.DeleteDeal(id); err != nil {
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

// searchDeal godoc
// @Summary Поиск соглашений по имени
// @Description Поиск соглашений по имени
// @Tags deal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param searchQuery query string true "query"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/deal/search [get]
func (dr *dealRoutes) searchDeal(ctx *gin.Context) {
	query := ctx.Query("searchQuery")
	deals, err := dr.s.SearchDeal(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponseWithData{
			Status:  -1,
			Message: "Not found",
			Data:    deals,
		})
		return
	}
	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    deals,
	})
}
