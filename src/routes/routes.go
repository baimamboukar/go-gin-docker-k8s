package routes

import "github.com/gin-gonic/gin"

func setupRoutes() *gin.RouterGroup {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	return versionRouter
}
