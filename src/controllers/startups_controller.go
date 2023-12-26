package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/baimamboukar/go-gin-docker-k8s/src/models"
)

func GetAllStartups(c *gin.Context) {
	startups, err := models.FetchAllStartups()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Startups fetched successfully", "status": "success", "data": startups})
}

func GetStartupByID(c *gin.Context) {
	startupID := c.Param("id")
	if startupID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
		return
	}
	startup, err := models.FetchStartup(startupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Startup fetched successfully", "status": "success", "data": startup})
}

func CreateStartup(c *gin.Context) {
	var input models.Startup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}
	savedStartup, err := input.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Startup saved successfully", "data": savedStartup})
}

func UpdateStartup(c *gin.Context) {
	startupID := c.Param("id")

	var updatedStartup *models.Startup
	if err := c.ShouldBindJSON(&updatedStartup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	updatedStartup, err := updatedStartup.UpdateStartup(startupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Startup updated successfully", "data": updatedStartup})
}

func DeleteStartup(c *gin.Context) {
	startupID := c.Param("id")
	if startupID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
		return
	}

	err := models.DeleteStartup(startupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Startup deleted successfully", "status": "success", "data": nil})
}
