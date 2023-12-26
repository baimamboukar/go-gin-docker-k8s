package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/baimamboukar/go-gin-docker-k8s/src/models"
	"github.com/baimamboukar/go-gin-docker-k8s/src/services"
)

func GetAllCompanies(c *gin.Context) {
	companies := services.GetAllCompanies()

	c.JSON(200, gin.H{"status": "success", "message": "Get companies success", "data": companies})
}

func GetCompanyByID(c *gin.Context) {
	// Implement logic to get a company by ID
	companyID := c.Param("id")
	c.JSON(200, gin.H{"message": "Get company by ID", "id": companyID})
}

func CreateCompany(context *gin.Context) {
	var input models.Company
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// savedCompany, err := input.

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func UpdateCompany(c *gin.Context) {
	// Implement logic to update a company
	c.JSON(200, gin.H{"message": "Update a company"})
}

func DeleteCompany(c *gin.Context) {
	// Implement logic to delete a company by ID
	companyID := c.Param("id")
	c.JSON(200, gin.H{"message": "Delete company by ID", "id": companyID})
}
