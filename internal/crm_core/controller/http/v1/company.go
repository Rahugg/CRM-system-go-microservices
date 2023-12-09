package v1

import (
	_ "crm_system/internal/auth/entity"
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type companyRoutes struct {
	s *service.Service
}

func newCompanyRoutes(handler *gin.RouterGroup, s *service.Service, MW *middleware.Middleware) {
	r := &companyRoutes{s: s}

	companyHandler := handler.Group("/company")
	{
		//middleware for users
		companyHandler.Use(MW.DeserializeUser("admin", "manager"))
		companyHandler.Use(MW.MetricsHandler())

		companyHandler.GET("/", r.getCompanies)
		companyHandler.GET("/:id", r.getCompany)
		companyHandler.POST("/", r.createCompany)
		companyHandler.PUT("/:id", r.updateCompany)
		companyHandler.DELETE("/:id", r.deleteCompany)
		companyHandler.GET("/search", r.searchCompany)
	}
}

// getCompanies godoc
// @Summary Получить список компаний
// @Description Получить список компаний
// @Tags company
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param sortBy query string false "sortBy"
// @Param sortOrder query string false "sortOrder"
// @Param phone query string false "filter by phone"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 404 {object} entity.CustomResponse
// @Router /v1/company/ [get]
func (cr *companyRoutes) getCompanies(ctx *gin.Context) {
	sortBy := ctx.Query("sortBy")
	sortOrder := ctx.Query("sortOrder")
	phone := ctx.Query("phone")

	companies, err := cr.s.GetCompanies(sortBy, sortOrder, phone)

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
		Data:    companies,
	})
}

// getCompany godoc
// @Summary Получить компанию по id
// @Description Получить компанию по id
// @Tags company
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id company"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/company/{id} [get]
func (cr *companyRoutes) getCompany(ctx *gin.Context) {
	id := ctx.Param("id")

	company, err := cr.s.GetCompany(id)

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
		Data:    company,
	})
}

// createCompany godoc
// @Summary Создать Компанию
// @Description Создать Компанию
// @Tags company
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param companyInput body entity.Company true "Create Company"
// @Success 201 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/company/ [post]
func (cr *companyRoutes) createCompany(ctx *gin.Context) {
	var company entity.Company

	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := cr.s.CreateCompany(company); err != nil {
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

// updateCompany godoc
// @Summary Редактировать компанию по id
// @Description Редактировать компанию по id
// @Tags company
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id company"
// @Param inputCompany body entity.NewCompany true "Update Company"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/company/{id} [put]
func (cr *companyRoutes) updateCompany(ctx *gin.Context) {
	id := ctx.Param("id")

	var newCompany entity.NewCompany

	if err := ctx.ShouldBindJSON(&newCompany); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := cr.s.UpdateCompany(newCompany, id); err != nil {
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

// deleteCompany godoc
// @Summary Удалить компанию по id
// @Description Удалить компанию по id
// @Tags company
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id company"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/company/{id} [delete]
func (cr *companyRoutes) deleteCompany(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := cr.s.DeleteCompany(id); err != nil {
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

// searchCompany godoc
// @Summary Поиск компании по имени
// @Description Поиск компании по имени
// @Tags company
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param searchQuery query string true "query"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/company/search [get]
func (cr *companyRoutes) searchCompany(ctx *gin.Context) {
	query := ctx.Query("searchQuery")
	companies, err := cr.s.SearchCompany(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponseWithData{
			Status:  -1,
			Message: "Not found",
			Data:    companies,
		})
		return
	}
	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    companies,
	})
}
