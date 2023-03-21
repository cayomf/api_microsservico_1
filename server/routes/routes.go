package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/controllers"
	"github.com/hyperyuri/webapi-with-go/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}
		vacancies := main.Group("vacancies", middlewares.Auth())
		{
			vacancies.GET("/", controllers.ShowAllVacancies)
			vacancies.GET("/:id", controllers.ShowVacancy)
			vacancies.POST("/", controllers.CreateVacancy)
			vacancies.PUT("/", controllers.EditVacancy)
			vacancies.DELETE("/:id", controllers.DeleteVacancy)
			vacancies.POST("/apply", controllers.ApplyForVacancy)
		}
		ads := main.Group("ads")
		{
			ads.GET("/last/:quantity", controllers.ShowLastVacancies)
		}

		main.POST("login", controllers.Login)
	}

	return router
}
