package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CarModelPropertyHandler struct {
	service *services.CarModelPropertyService
}

func NewCarModelPropertyHandler(cfg *config.Config) *CarModelPropertyHandler {
	return &CarModelPropertyHandler{
		service: services.NewCarModelPropertyService(cfg),
	}
}

func (h *CarModelPropertyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CarModelPropertyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CarModelPropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CarModelPropertyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CarModelPropertyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
