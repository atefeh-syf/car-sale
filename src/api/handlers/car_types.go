package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CarTypeHandler struct {
	service *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{
		service: services.NewCarTypeService(cfg),
	}
}

func (h *CarTypeHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CarTypeHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CarTypeHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CarTypeHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CarTypeHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
