package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type taskRoutes struct {
	s *service.Service
}

func newTaskRoutes(handler *gin.RouterGroup, s *service.Service, MW *middleware.Middleware) {
	r := &taskRoutes{s: s}

	taskHandler := handler.Group("/task")
	{
		//middleware for users
		taskHandler.Use(MW.MetricsHandler())
		taskHandler.GET("/deal/:dealId", MW.DeserializeUser("any"), r.getTasks)
		taskHandler.GET("/:id", MW.DeserializeUser("any"), r.getTask)
		taskHandler.POST("/", MW.DeserializeUser("manager", "admin"), r.createTask)
		taskHandler.POST("/vote", MW.DeserializeUser("any"), r.vote)
		taskHandler.PUT("/:id", MW.DeserializeUser("manager", "admin"), r.updateTask)
		taskHandler.DELETE("/:id", MW.DeserializeUser("manager", "admin"), r.deleteTask)
		taskHandler.GET("/changes/:id", MW.DeserializeUser("manager", "admin"), r.getChangesOfTask)
		taskHandler.GET("/search", r.searchTask)
	}
}

// getTasks godoc
// @Summary Получить список заданий
// @Description Получить список заданий
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param sortBy query string false "sortBy"
// @Param sortOrder query string false "sortOrder"
// @Param state query string false "filter by state"
// @Param dealId path string true "dealId"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 404 {object} entity.CustomResponse
// @Router /v1/task/deal/{dealId} [get]
func (tr *taskRoutes) getTasks(ctx *gin.Context) {
	dealId := ctx.Param("dealId")
	sortBy := ctx.Query("sortBy")
	sortOrder := ctx.Query("sortOrder")
	state := ctx.Query("state")

	user, exists := ctx.MustGet("currentUser").(*entity.User)
	if !exists {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponse{
			Status:  -1,
			Message: fmt.Errorf("the current user did not authorize or does not exist").Error(),
		})
		return
	}

	tasks, err := tr.s.GetTasks(dealId, sortBy, sortOrder, state, user)

	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponse{
			Status:  -2,
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

// getTask godoc
// @Summary Получить задание по id
// @Description Получить задание по id
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id task"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/task/{id} [get]
func (tr *taskRoutes) getTask(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := tr.s.GetTask(id)

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

// createTask godoc
// @Summary Создать Задание
// @Description Создать Задание
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param taskInput body entity.TaskInput true "Create task"
// @Success 201 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/task/ [post]
func (tr *taskRoutes) createTask(ctx *gin.Context) {
	var task entity.TaskInput

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if time.Now().After(task.DueDate) {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: "Due date must be in the future",
		})
		return
	}

	if err := tr.s.CreateTask(&task); err != nil {
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

// vote godoc
// @Summary Проголосовать за задание
// @Description Проголосовать за задание
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param voteInput body entity.VoteInput true "Create Vote"
// @Success 201 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/task/vote [post]
func (tr *taskRoutes) vote(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(*entity.User)
	var voteInput *entity.VoteInput
	if err := ctx.ShouldBindJSON(&voteInput); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	err := tr.s.Vote(user, voteInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
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

// updateTask godoc
// @Summary Редактировать задание по id
// @Description Редактировать задание по id
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id task"
// @Param inputTask body entity.TaskEditInput true "Update task"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/task/{id} [put]
func (tr *taskRoutes) updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	user := ctx.MustGet("currentUser").(*entity.User)

	var task entity.TaskEditInput

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := tr.s.UpdateTask(task, id, user); err != nil {
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

// deleteTask godoc
// @Summary Удалить задание по id
// @Description Удалить задание по id
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id task"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/task/{id} [delete]
func (tr *taskRoutes) deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := tr.s.DeleteTask(id); err != nil {
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

// getChangesOfTask godoc
// @Summary Получить историю изменений по заданию
// @Description Получить историю изменений по заданию
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id task"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/task/changes/{id} [get]
func (tr *taskRoutes) getChangesOfTask(ctx *gin.Context) {
	id := ctx.Param("id")
	todoChanges, err := tr.s.GetChangesOfTask(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    todoChanges,
	})

}

// searchTask godoc
// @Summary Поиск задание по имени
// @Description Поиск задание по имени
// @Tags task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param searchQuery query string true "query"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/task/search [get]
func (tr *taskRoutes) searchTask(ctx *gin.Context) {
	query := ctx.Query("searchQuery")
	tasks, err := tr.s.SearchTask(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponseWithData{
			Status:  -1,
			Message: "Not found",
			Data:    tasks,
		})
		return
	}
	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    tasks,
	})
}
