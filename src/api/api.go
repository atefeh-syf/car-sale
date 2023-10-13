package api

import (
	"fmt"

	"github.com/atefeh-syf/car-sale/api/middlewares"
	"github.com/atefeh-syf/car-sale/api/routers"
	validation "github.com/atefeh-syf/car-sale/api/validations"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	//r1 := gin.Default()

	RegisterValidator()

	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger()/*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}
}
