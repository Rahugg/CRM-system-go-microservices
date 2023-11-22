package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/service"
	"crm_system/pkg/crm_core/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type dealRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newDealRoutes(handler *gin.RouterGroup, s *service.Service, MW *middleware.Middleware) {
	r := &dealRoutes{s: s}

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
	deals, err := dr.s.GetDeals(ctx)

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
func (dr *dealRoutes) getDeal(ctx *gin.Context) {
	id := ctx.Param("id")

	deal, err := dr.s.GetDeal(ctx, id)

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
func (dr *dealRoutes) createDeal(ctx *gin.Context) {
	var deal entity.Deal

	if err := ctx.ShouldBindJSON(&deal); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := dr.s.CreateDeal(ctx, deal); err != nil {
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

	if err := dr.s.UpdateDeal(ctx, newDeal, id); err != nil {
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
func (dr *dealRoutes) deleteDeal(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := dr.s.DeleteDeal(ctx, id); err != nil {
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
