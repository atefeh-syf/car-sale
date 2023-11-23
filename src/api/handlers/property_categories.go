package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{
		service: services.NewPropertyCategoryService(cfg),
	}
}

func (h *PropertyCategoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *PropertyCategoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *PropertyCategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *PropertyCategoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *PropertyCategoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
