package handlers

import (
	"fmt"
	"net/http"

	"github.com/atefeh-syf/car-sale/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{

}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK,  helper.GenerateBaseResponse("working!", true, 0))
	return
}

func (h *HealthHandler) HealthPost(c *gin.Context){
	c.JSON(http.StatusOK,   helper.GenerateBaseResponse("working! post", true, 0))
	return
}

func (h *HealthHandler) HealthPostById(c *gin.Context){
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("Working post by id: %s!", id))
	return
}