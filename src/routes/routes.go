package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	versionRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))
	startupsGroupRouter(versionRouter)
	companiesGroupRouter(versionRouter)

	return r
}
