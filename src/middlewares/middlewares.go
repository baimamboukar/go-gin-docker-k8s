package middlewares

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterMiddlewares(router *gin.Engine) {
	logger, _ := zap.NewProduction()
	//router.Use(AuthMiddleware())
	router.Use(LoggerMiddleware())
	router.Use(PrometheusMiddleware())
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
}
