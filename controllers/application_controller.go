package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/models"
	"github.com/hyperyuri/webapi-with-go/services"
)

type ApplicationController interface {
	ApplyForVacancy(c *gin.Context)
}

type applicationControllerImpl struct {
	applicationService services.ApplicationService
}

func NewApplicationController(applicationService services.ApplicationService) ApplicationController {
	return &applicationControllerImpl{applicationService: applicationService}
}

func (a *applicationControllerImpl) ApplyForVacancy(c *gin.Context) {
	var application models.Application
	err := c.ShouldBindJSON(&application)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = a.applicationService.ApplyForVacancy(&application)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot apply for vacancy: " + err.Error(),
		})
		return
	}

	c.JSON(200, application)
}
