package handlers

import (
	"net/http"

	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/api/helper"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/services"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config)  *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{
		service: service,
	}
}

func (h *UsersHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	err = h.service.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err) )
			return
	}
	// call internal sms service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}