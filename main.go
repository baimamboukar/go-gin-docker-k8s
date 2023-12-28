package main

import (
	"github.com/baimamboukar/go-gin-docker-k8s/src/middlewares"
	"github.com/baimamboukar/go-gin-docker-k8s/src/models"
	"github.com/baimamboukar/go-gin-docker-k8s/src/routes"
	"github.com/baimamboukar/go-gin-docker-k8s/src/utils"
)

func main() {
	utils.LoadEnv()
	router := routes.SetupRoutes()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	middlewares.RegisterMiddlewares(router)
	router.Run(":8113")

}
