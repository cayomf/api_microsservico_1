package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/controllers"
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/repositories"
	"github.com/hyperyuri/webapi-with-go/server/middlewares"
	"github.com/hyperyuri/webapi-with-go/services"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	db := database.GetDatabase()

	jwtService := services.NewJWTService()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, jwtService)
	userController := controllers.NewUserController(userService)

	vacancyRepository := repositories.NewVacancyRepository(db)
	vacancyService := services.NewVacancyService(vacancyRepository)
	vacancyController := controllers.NewVacancyController(vacancyService)

	applicationRepository := repositories.NewApplicationRepository(db)
	applicationService := services.NewApplicationService(applicationRepository)
	applicationController := controllers.NewApplicationController(applicationService)

	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", userController.CreateUser)
		}
		vacancies := main.Group("vacancies", middlewares.Auth())
		{
			vacancies.GET("/", vacancyController.ShowAllVacancies)
			vacancies.GET("/:id", vacancyController.ShowVacancy)
			vacancies.POST("/", vacancyController.CreateVacancy)
			vacancies.PUT("/", vacancyController.EditVacancy)
			vacancies.DELETE("/:id", vacancyController.DeleteVacancy)
			vacancies.POST("/apply", applicationController.ApplyForVacancy)
		}
		ads := main.Group("ads")
		{
			ads.GET("/last/:quantity", vacancyController.ShowLastVacancies)
		}

		main.POST("login", userController.Login)
	}

	router.GET("metrics", prometheusHandler())

	return router
}

func prometheusHandler() gin.HandlerFunc {
	promHandler := promhttp.Handler()

	return func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	}
}
