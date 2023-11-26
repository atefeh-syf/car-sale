package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CarModelImageHandler struct {
	service *services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{
		service: services.NewCarModelImageService(cfg),
	}
}

func (h *CarModelImageHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CarModelImageHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CarModelImageHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CarModelImageHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CarModelImageHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
