package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/baimamboukar/go-gin-docker-k8s/src/controllers"
)

func startupsGroupRouter(baseRouter *gin.RouterGroup) {
	startups := baseRouter.Group("/startups")

	startups.GET("/all", controllers.GetAllStartups)
	startups.GET("/get/:id", controllers.GetStartupByID)
	startups.POST("/create", controllers.CreateStartup)
	startups.PATCH("/update", controllers.UpdateStartup)
	startups.PUT("/update", controllers.UpdateStartup)
	startups.DELETE("/delete/:id", controllers.DeleteStartup)
}
