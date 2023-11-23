package v1

import (
	"crm_system/internal/auth/controller/http/middleware"
	"crm_system/internal/auth/entity"
	"crm_system/internal/auth/service"
	"crm_system/pkg/auth/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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
		userHandler.POST("/login", r.signIn)
		userHandler.GET("/logout", r.logout)
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
	ur.signUp(ctx, 1, true, "admin")
}

// TODO: change the verified to an email with kafka
func (ur *userRoutes) signUpManager(ctx *gin.Context) {
	verified := ur.s.Config.Gin.Mode == "debug"
	ur.signUp(ctx, 2, verified, "manager")
}

// TODO: change the verified to an email with kafka
func (ur *userRoutes) signUpSalesRep(ctx *gin.Context) {
	verified := ur.s.Config.Gin.Mode == "debug"
	ur.signUp(ctx, 3, verified, "sales_rep")
}

// TODO: change the verified to an email with kafka
func (ur *userRoutes) signUpSupportRep(ctx *gin.Context) {
	verified := ur.s.Config.Gin.Mode == "debug"
	ur.signUp(ctx, 4, verified, "support_rep")
}

func (ur *userRoutes) signUp(ctx *gin.Context, roleId uint, verified bool, provider string) {
	var payload entity.SignUpInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	data, err := ur.s.SignUp(ctx, &payload, roleId, verified, provider)
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
func (ur *userRoutes) signIn(ctx *gin.Context) {
	var payload *entity.SignInInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	data, err := ur.s.SignIn(ctx, payload)
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

func (ur *userRoutes) logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}

func (ur *userRoutes) getUsers(ctx *gin.Context) {
	users, err := ur.s.GetUsers(ctx)

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

func (ur *userRoutes) getUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := ur.s.GetUser(ctx, id)

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

	if err := ur.s.UpdateUser(ctx, newUser, id); err != nil {
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
func (ur *userRoutes) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := ur.s.DeleteUser(ctx, id); err != nil {
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
func (ur *userRoutes) createUser(ctx *gin.Context) {

	var user *entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := ur.s.CreateUser(ctx, user); err != nil {
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
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}

func (ur *userRoutes) searchUser(ctx *gin.Context) {
	query := ctx.Query("query")
	users, err := ur.s.SearchUser(ctx, query)
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
