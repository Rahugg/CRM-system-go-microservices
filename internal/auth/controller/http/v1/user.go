package v1

import (
	"crm_system/internal/auth/controller/http/middleware"
	"crm_system/internal/auth/entity"
	"crm_system/internal/auth/service"
	_ "crm_system/internal/crm_core/entity"
	"crm_system/pkg/auth/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type userRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newUserRoutes(handler *gin.RouterGroup, s *service.Service, MW *middleware.Middleware) {
	r := &userRoutes{s: s}

	adminHandler := handler.Group("/admin/user")
	{
		adminHandler.Use(MW.CustomLogger())
		adminHandler.Use(MW.DeserializeUser("admin"))

		adminHandler.POST("/register", r.signUpAdmin)
		adminHandler.GET("/", r.getUsers)
		adminHandler.GET("/:id", r.getUser)
		adminHandler.PUT("/:id", r.updateUser)
		adminHandler.DELETE("/:id", r.deleteUser)
		adminHandler.POST("/", r.createUser)
		adminHandler.GET("/search", r.searchUser)

		adminHandler.GET("/test", r.test)
	}

	userHandler := handler.Group("/user")
	{
		userHandler.POST("/register", r.signUpManager)
		userHandler.POST("/confirm", r.confirmUser)
		userHandler.POST("/login", r.signIn)
		userHandler.POST("/register/sales", MW.DeserializeUser("manager"), r.signUpSalesRep)
		userHandler.POST("/register/support", MW.DeserializeUser("manager"), r.signUpSupportRep)

		userHandler.GET("/me", MW.DeserializeUser("admin", "manager", "sales_rep", "support_rep"), r.getMe)
		userHandler.PUT("/me", MW.DeserializeUser("admin", "manager", "sales_rep", "support_rep"), r.updateMe)
	}
}

func (ur *userRoutes) test(ctx *gin.Context) {
	ur.l.Info("hello from controller")
	ctx.JSON(http.StatusOK, "test")
}

func (ur *userRoutes) signUpAdmin(ctx *gin.Context) {
	ur.signUp(ctx, 1, "admin")
}

func (ur *userRoutes) signUpManager(ctx *gin.Context) {
	ur.signUp(ctx, 2, "manager")
}

func (ur *userRoutes) signUpSalesRep(ctx *gin.Context) {
	ur.signUp(ctx, 3, "sales_rep")
}

func (ur *userRoutes) signUpSupportRep(ctx *gin.Context) {
	ur.signUp(ctx, 4, "support_rep")
}

// signUp godoc
// @Summary Зарегистрировать пользователя с помощью signUp
// @Description Зарегистрировать пользователя с помощью signUp
// @Tags auth
// @Accept json
// @Produce json
// @Param signUp body entity.SignUpInput true "Register user"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/user/register [post]
func (ur *userRoutes) signUp(ctx *gin.Context, roleId uint, provider string) {
	var payload entity.SignUpInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	id, err := ur.s.SignUp(&payload, roleId, provider)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go ur.createUserCode(ctx, id, &wg)
	wg.Wait()

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "go confirm yourself",
	})
}

// signIn godoc
// @Summary Авторизовать пользователя с помощью signUp
// @Description Авторизовать пользователя с помощью signIn
// @Tags auth
// @Accept json
// @Produce json
// @Param signIn body entity.SignInInput true "Authorize user"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/user/login [post]
func (ur *userRoutes) signIn(ctx *gin.Context) {
	var payload *entity.SignInInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	data, err := ur.s.SignIn(payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    data,
	})
}

// getUsers godoc
// @Summary Получить список пользователей
// @Description Получить список пользователей
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param sortBy query string false "sortBy"
// @Param sortOrder query string false "sortOrder"
// @Param age query string false "filter by age"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/admin/user/ [get]
func (ur *userRoutes) getUsers(ctx *gin.Context) {
	sortBy := ctx.Query("sortBy")
	sortOrder := ctx.Query("sortOrder")
	age := ctx.Query("age")

	users, err := ur.s.GetUsers(sortBy, sortOrder, age)

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
		Data:    users,
	})
}

// getUser godoc
// @Summary Получить пользователя по id
// @Description Получить пользователя по id
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id user"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/admin/user/{id} [get]
func (ur *userRoutes) getUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := ur.s.GetUser(id)

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
		Data:    user,
	})
}

// updateUser godoc
// @Summary Редактировать пользователя по id
// @Description Редактировать пользователя по id
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id user"
// @Param inputUser body entity.User true "Update User"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/admin/user/{id} [put]
func (ur *userRoutes) updateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var newUser *entity.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := ur.s.UpdateUser(newUser, id); err != nil {
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

// deleteUser godoc
// @Summary Удалить пользователя по id
// @Description Удалить пользователя по id
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id user"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/admin/user/{id} [delete]
func (ur *userRoutes) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := ur.s.DeleteUser(id); err != nil {
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

// createUser godoc
// @Summary Создать Пользователя
// @Description Создать Пользователя
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param inputUser body entity.User true "Create User"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/admin/user/ [post]
func (ur *userRoutes) createUser(ctx *gin.Context) {
	var user *entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := ur.s.CreateUser(user); err != nil {
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

// getMe godoc
// @Summary Получить информацию о себе
// @Description Получить информацию о себе
// @Tags cabinet
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/user/me/ [get]
func (ur *userRoutes) getMe(ctx *gin.Context) {
	user, err := ur.s.GetMe(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    user,
	})
}

// updateMe godoc
// @Summary Поменять информацию о себе
// @Description Поменять информацию о себе
// @Tags cabinet
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body entity.User true "Поменять информацию о себе"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/user/me/ [put]
func (ur *userRoutes) updateMe(ctx *gin.Context) {
	var newUser *entity.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := ur.s.UpdateMe(ctx, newUser); err != nil {
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

// searchUser godoc
// @Summary Поиск пользователя по имени
// @Description Поиск пользователя по имени
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param searchQuery query string true "query"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/admin/user/search [get]
func (ur *userRoutes) searchUser(ctx *gin.Context) {
	query := ctx.Query("searchQuery")
	users, err := ur.s.SearchUser(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponseWithData{
			Status:  -1,
			Message: "Not found",
			Data:    users,
		})
		return
	}
	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    users,
	})
}

func (ur *userRoutes) createUserCode(ctx *gin.Context, id string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := ur.s.CreateUserCode(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}
}

// confirmUser godoc
// @Summary Подтвердить пользователя по коду
// @Description Подтвердить пользователя по коду
// @Tags auth
// @Accept json
// @Produce json
// @Param inputCode body entity.InputCode true "Код"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/user/confirm/ [post]
func (ur *userRoutes) confirmUser(ctx *gin.Context) {
	var code entity.InputCode

	err := ctx.ShouldBindJSON(&code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	err = ur.s.ConfirmUser(code.Code)
	if err != nil {
		ur.l.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "User confirmed",
	})
}
