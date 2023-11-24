package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type ColorHandler struct {
	service *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{
		service: services.NewColorService(cfg),
	}
}

func (h *ColorHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *ColorHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *ColorHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *ColorHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *ColorHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
