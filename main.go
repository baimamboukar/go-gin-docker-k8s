package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")

	companies := versionRouter.Group("/companies")
	companies.GET("/all", GetCompanies)

}

func GetCompanies(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Got all companies"})
}
