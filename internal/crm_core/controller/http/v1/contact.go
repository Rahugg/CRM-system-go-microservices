package v1

import (
	"crm_system/internal/crm_core/controller/http/middleware"
	"crm_system/internal/crm_core/entity"
	"crm_system/internal/crm_core/service"
	"crm_system/pkg/crm_core/cache"
	"crm_system/pkg/crm_core/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type contactRoutes struct {
	s            *service.Service
	l            *logger.Logger
	contactCache cache.Contact
}

func newContactRoutes(handler *gin.RouterGroup, s *service.Service, MW *middleware.Middleware, cc cache.Contact) {
	r := &contactRoutes{s: s, contactCache: cc}

	contactHandler := handler.Group("/contact")
	{
		//middleware for users
		contactHandler.Use(MW.MetricsHandler())
		contactHandler.Use(MW.DeserializeUser("manager", "admin"))

		contactHandler.GET("/", r.getContacts)
		contactHandler.GET("/:id", r.getContact)
		contactHandler.POST("/", r.createContact)
		contactHandler.PUT("/:id", r.updateContact)
		contactHandler.DELETE("/:id", r.deleteContact)
		contactHandler.GET("/search", r.searchContact)
	}
}

// getContacts godoc
// @Summary Получить список контактов
// @Description Получить список контактов
// @Tags contact
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param sortBy query string false "sortBy"
// @Param sortOrder query string false "sortOrder"
// @Param phone query string false "filter by phone"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 404 {object} entity.CustomResponse
// @Router /v1/contact/ [get]
func (cr *contactRoutes) getContacts(ctx *gin.Context) {
	sortBy := ctx.Query("sortBy")
	sortOrder := ctx.Query("sortOrder")
	phone := ctx.Query("phone")

	contacts, err := cr.s.GetContacts(sortBy, sortOrder, phone)

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
		Data:    contacts,
	})
}

// getContact godoc
// @Summary Получить контакт по id
// @Description Получить контакт по id
// @Tags contact
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id contact"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/contact/{id} [get]
func (cr *contactRoutes) getContact(ctx *gin.Context) {
	id := ctx.Param("id")

	contact, err := cr.contactCache.Get(ctx, id)
	if err != nil {
		return
	}

	if contact == nil {
		contact, err = cr.s.GetContact(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, &entity.CustomResponse{
				Status:  -2,
				Message: err.Error(),
			})
			return
		}

		err = cr.contactCache.Set(ctx, id, contact)
		if err != nil {
			log.Printf("could not cache contact with id %s: %v", id, err)
		}
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  -1,
		Message: "OK",
		Data:    contact,
	})
}

// createContact godoc
// @Summary Создать Контакт
// @Description Создать Контакт
// @Tags contact
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param contactInput body entity.Contact true "Create Contact"
// @Success 201 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/contact/ [post]
func (cr *contactRoutes) createContact(ctx *gin.Context) {
	var contact entity.Contact

	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := cr.s.CreateContact(contact); err != nil {
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

// updateContact godoc
// @Summary Редактировать контакт по id
// @Description Редактировать контакт по id
// @Tags contact
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id contact"
// @Param inputContact body entity.Contact true "Update Contact"
// @Success 200 {object} entity.CustomResponse
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/contact/{id} [put]
func (cr *contactRoutes) updateContact(ctx *gin.Context) {
	id := ctx.Param("id")

	var newContact entity.Contact

	if err := ctx.ShouldBindJSON(&newContact); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := cr.s.UpdateContact(newContact, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}

	err := cr.contactCache.Set(ctx, id, &newContact)
	if err != nil {
		cr.l.Debug("could not cache contact with id %s: %v", id, err)
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}

// deleteContact godoc
// @Summary Удалить контакт по id
// @Description Удалить контакт по id
// @Tags contact
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id contact"
// @Success 200 {object} entity.CustomResponse
// @Failure 204 {object} entity.CustomResponse
// @Router /v1/contact/{id} [delete]
func (cr *contactRoutes) deleteContact(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := cr.s.DeleteContact(id); err != nil {
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

// searchContact godoc
// @Summary Поиск контакта по имени
// @Description Поиск контакта по имени
// @Tags contact
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param searchQuery query string true "query"
// @Success 200 {object} entity.CustomResponseWithData
// @Failure 400 {object} entity.CustomResponse
// @Router /v1/contact/search [get]
func (cr *contactRoutes) searchContact(ctx *gin.Context) {
	query := ctx.Query("searchQuery")
	contacts, err := cr.s.SearchContact(query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponseWithData{
			Status:  -1,
			Message: "Not found",
			Data:    contacts,
		})
		return
	}
	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    contacts,
	})
}
