package routers

import (
	"github.com/atefeh-syf/car-sale/api/handlers"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/gin-gonic/gin"
)

func Country(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountryHandler(cfg)

	router.POST("/", h.Create)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
	router.GET("/:id", h.GetById)
	router.POST("/get-by-filter", h.GetByFilter)

}

func City(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCityHandler(cfg)

	router.POST("/", h.Create)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
	router.GET("/:id", h.GetById)
	router.POST("/get-by-filter", h.GetByFilter)
}

func File(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
