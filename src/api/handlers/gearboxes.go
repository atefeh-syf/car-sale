package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type GearboxHandler struct {
	service *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{
		service: services.NewGearboxService(cfg),
	}
}

func (h *GearboxHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *GearboxHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *GearboxHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *GearboxHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *GearboxHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
