package routes

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	startupsGroupRouter(versionRouter)
	companiesGroupRouter(versionRouter)

	return r
}
