package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/baimamboukar/go-gin-docker-k8s/src/models"
)

func GetAllCompanies(c *gin.Context) {
	companies, err := models.FetchAllCompanies()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Companies fetched successfully", "status": "success", "data": companies})
}

func GetCompanyByID(c *gin.Context) {
	companyID := c.Param("id")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Company ID is required"})
		return
	}
	company, err := models.FetchCompany(companyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company fetched successfully", "status": "success", "data": company})
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

	context.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Comnapy saved successfuly", "data": savedCompany})
}
func UpdateCompany(c *gin.Context) {
	// Get the company ID from the URL parameter
	companyID := c.Param("id")

	// Bind the updated data from the request
	var updatedCompany *models.Company
	if err := c.ShouldBindJSON(&updatedCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	updatedCompany, err := updatedCompany.UpdateCompany(companyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Company updated successfully", "data": updatedCompany})
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
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully", "status": "success", "data": nil})

}
