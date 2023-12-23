package routes

import (
	"github.com/gin-gonic/gin"

	service "github.com/baimamboukar/go-gin-docker-k8s/src/services"
)

func companiesGroupRouter(baseRouter *gin.RouterGroup) {

	companies := baseRouter.Group("/companies")
	companies.GET("/all", GetAllCompanies)
	companies.GET("/:id", GetCompanyByID)
	companies.POST("/create", CreateCompany)
	companies.PATCH("/update", UpdateCompany)
	companies.PUT("/update", UpdateCompany)
	companies.DELETE("/:id", DeleteCompany)
}

func GetAllCompanies(c *gin.Context) {
	companies := service.GetAllCompanies()

	c.JSON(200, gin.H{"status": "success", "message": "Get companies success", "data": companies})
}

func GetCompanyByID(c *gin.Context) {
	// Implement logic to get a company by ID
	companyID := c.Param("id")
	c.JSON(200, gin.H{"message": "Get company by ID", "id": companyID})
}

func CreateCompany(c *gin.Context) {
	// Implement logic to create a new company
	c.JSON(200, gin.H{"message": "Create a new company"})
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
