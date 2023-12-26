package main

import (
	"github.com/baimamboukar/go-gin-docker-k8s/src/middlewares"
	"github.com/baimamboukar/go-gin-docker-k8s/src/models"
	"github.com/baimamboukar/go-gin-docker-k8s/src/routes"
)

func main() {
	router := routes.SetupRoutes()
	models.OpenDatabaseConnection()
	middlewares.RegisterMiddlewares(router)
	router.Run(":8081")

}
