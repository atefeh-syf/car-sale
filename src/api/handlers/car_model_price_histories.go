package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CarModelPriceHistoryHandler struct {
	service *services.CarModelPriceHistoryService
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	return &CarModelPriceHistoryHandler{
		service: services.NewCarModelPriceHistoryService(cfg),
	}
}

func (h *CarModelPriceHistoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CarModelPriceHistoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CarModelPriceHistoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CarModelPriceHistoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CarModelPriceHistoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
