package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CarModelColorHandler struct {
	service *services.CarModelColorService
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	return &CarModelColorHandler{
		service: services.NewCarModelColorService(cfg),
	}
}

func (h *CarModelColorHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CarModelColorHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CarModelColorHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CarModelColorHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CarModelColorHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
