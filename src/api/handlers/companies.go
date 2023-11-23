package handlers

import (
	_ "github.com/atefeh-syf/car-sale/api/dto"
	_ "github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{
		service: services.NewCompanyService(cfg),
	}
}

func (h *CompanyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

func (h *CompanyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

func (h *CompanyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

func (h *CompanyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

func (h *CompanyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFiler)
}
