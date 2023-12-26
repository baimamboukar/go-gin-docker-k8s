package middlewares

import "github.com/gin-gonic/gin"

func RegisterMiddlewares(router *gin.Engine) {
	router.Use(AuthMiddleware())
	router.Use(LoggerMiddleware())
}
