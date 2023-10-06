package api

import (
	"fmt"

	"github.com/atefeh-syf/car-sale/api/routers"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	//r1 := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}
