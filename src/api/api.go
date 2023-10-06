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

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}
