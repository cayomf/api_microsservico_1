package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/models"
)

func ApplyForVacancy(c *gin.Context) {
	db := database.GetDatabase()

	var application models.Application

	err := c.ShouldBindJSON(&application)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&application).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot apply for vacancy: " + err.Error(),
		})
		return
	}

	c.JSON(200, application)
}
