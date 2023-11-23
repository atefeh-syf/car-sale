package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{
		service: services.NewCountryService(cfg),
	}
}

func (h *CountryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CountryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CountryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CountryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CountryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
