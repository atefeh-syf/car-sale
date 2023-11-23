package api

import (
	"fmt"
	"github.com/atefeh-syf/car-sale/pkg/logging"

	"github.com/atefeh-syf/car-sale/api/middlewares"
	"github.com/atefeh-syf/car-sale/api/routers"
	validation "github.com/atefeh-syf/car-sale/api/validations"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	//r1 := gin.Default()

	RegisterValidators()

	r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler) /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.StartUp, err.Error(), nil)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.StartUp, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.StartUp, err.Error(), nil)
		}
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		users_router := v1.Group("/users")
		countries_router := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		cities_router := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		files_router := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		property_categories_router := v1.Group("/property-categories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		properties_router := v1.Group("/properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		routers.Health(health)
		routers.TestRouter(test_router)
		routers.User(users_router, cfg)
		routers.Country(countries_router, cfg)
		routers.City(cities_router, cfg)
		routers.File(files_router, cfg)
		routers.Property(properties_router, cfg)
		routers.PropertyCategory(property_categories_router, cfg)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang car sale api"
	docs.SwaggerInfo.Description = "golang car sale api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprint("localhost:%s", cfg.Server.InternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
