package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CarModelHandler struct {
	service *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{
		service: services.NewCarModelService(cfg),
	}
}

func (h *CarModelHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CarModelHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CarModelHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CarModelHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CarModelHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
