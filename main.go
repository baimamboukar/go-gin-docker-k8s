package main

import "github.com/baimamboukar/go-gin-docker-k8s/src/routes"

func main() {
	router := routes.SetupRoutes()
	router.Run(":8081")

}
