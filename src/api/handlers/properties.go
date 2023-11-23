package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	service *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	return &PropertyHandler{
		service: services.NewPropertyService(cfg),
	}
}

func (h *PropertyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *PropertyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *PropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *PropertyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *PropertyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
