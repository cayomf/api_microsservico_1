package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/models"
	"github.com/hyperyuri/webapi-with-go/services"
)

type UserController interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}

type userControllerImpl struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

func (u *userControllerImpl) CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = u.userService.CreateUser(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func (u *userControllerImpl) Login(c *gin.Context) {
	var p models.Login
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	token, err := u.userService.Login(p.Email, p.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
