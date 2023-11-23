package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		service: services.NewCityService(cfg),
	}
}

func (h *CityHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CityHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CityHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CityHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CityHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
