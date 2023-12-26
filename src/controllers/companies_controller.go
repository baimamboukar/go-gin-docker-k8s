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
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}
	savedCompany, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": "failed", "message": "Comnapy saved successfuly", "data": savedCompany})
}

func UpdateCompany(c *gin.Context) {
}

func DeleteCompany(c *gin.Context) {
	// Implement logic to delete a company
	companyID := c.Param("id")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Company ID is required"})
		return
	}
	err := models.DeleteCompany(companyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})

}
