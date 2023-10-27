package routers

import (
	"github.com/atefeh-syf/car-sale/api/handlers"
	"github.com/atefeh-syf/car-sale/api/middlewares"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp",  middlewares.OtpLimiter(cfg) ,h.SendOtp)
}