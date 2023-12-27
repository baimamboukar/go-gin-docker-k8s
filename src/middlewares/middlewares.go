package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterMiddlewares(router *gin.Engine) {
	router.Use(AuthMiddleware())
	router.Use(LoggerMiddleware())
	fmt.Println("---Registrating Prometheus----")
	router.Use(PrometheusMiddleware())
	fmt.Println("---Prometheus Registrated----")
}
