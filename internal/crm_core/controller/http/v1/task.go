package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/service"
	"crm_system/pkg/crm_core/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type taskRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newTaskRoutes(handler *gin.RouterGroup, s *service.Service, MW *middleware.Middleware) {
	r := &taskRoutes{s: s}

	taskHandler := handler.Group("/task")
	{
		//middleware for users
		taskHandler.GET("/", r.getTasks)
		taskHandler.GET("/:id", r.getTask)
		taskHandler.POST("/", r.createTask)
		taskHandler.PUT("/:id", r.updateTask)
		taskHandler.DELETE("/:id", r.deleteTask)
	}
}

func (tr *taskRoutes) getTasks(ctx *gin.Context) {
	tasks, err := tr.s.GetTasks(ctx)

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
		Data:    tasks,
	})
}
func (tr *taskRoutes) getTask(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := tr.s.GetTask(ctx, id)

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
		Data:    task,
	})
}
func (tr *taskRoutes) createTask(ctx *gin.Context) {
	var task entity.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := tr.s.CreateTask(ctx, task); err != nil {
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
func (tr *taskRoutes) updateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var task entity.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := tr.s.UpdateTask(ctx, task, id); err != nil {
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
func (tr *taskRoutes) deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := tr.s.DeleteTask(ctx, id); err != nil {
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
