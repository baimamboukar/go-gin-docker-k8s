package routes

import "github.com/gin-gonic/gin"

func setupRoutes() {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	startupsGroupRouter(versionRouter)
	companiesGroupRouter(versionRouter)

	r.Run(":8081")
}
