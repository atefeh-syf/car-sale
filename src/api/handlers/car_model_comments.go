package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CarModelCommentHandler struct {
	service *services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{
		service: services.NewCarModelCommentService(cfg),
	}
}

func (h *CarModelCommentHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CarModelCommentHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CarModelCommentHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CarModelCommentHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CarModelCommentHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
