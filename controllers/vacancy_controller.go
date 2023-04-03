package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/models"
	"github.com/hyperyuri/webapi-with-go/services"
	"github.com/prometheus/client_golang/prometheus"
)

type VacancyController struct {
	vacancyService services.VacancyService
}

var RequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "requests_total",
	Help: "Total number of HTTP requests.",
})

var RequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "request_duration_seconds",
	Help:    "Duração das requisições",
	Buckets: prometheus.LinearBuckets(0.1, 0.1, 10), // 10 buckets, cada um com intervalo de 0.1 segundos
})

func NewVacancyController(vacancyService services.VacancyService) *VacancyController {
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(RequestDuration)

	return &VacancyController{vacancyService: vacancyService}
}

func (vc *VacancyController) ShowAllVacancies(c *gin.Context) {
	vacancies, err := vc.vacancyService.GetAll()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find vacancies: " + err.Error(),
		})
		return
	}

	c.JSON(200, vacancies)
}

func (vc *VacancyController) ShowLastVacancies(c *gin.Context) {
	quantity, _ := strconv.ParseInt(c.Param("quantity"), 10, 64)

	vacancies, err := vc.vacancyService.GetLast(int(quantity))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot get last elements: " + err.Error(),
		})
		return
	}

	c.JSON(200, vacancies)
}

func (vc *VacancyController) ShowVacancy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	vacancy, err := vc.vacancyService.GetById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find vacancy by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, vacancy)
}

func (vc *VacancyController) CreateVacancy(c *gin.Context) {
	var vacancy models.Vacancy
	startTime := time.Now()

	err := c.ShouldBindJSON(&vacancy)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = vc.vacancyService.Create(&vacancy)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create vacancy: " + err.Error(),
		})
		return
	}

	duration := time.Since(startTime)
	durationSeconds := duration.Seconds()

	RequestsTotal.Inc()
	RequestDuration.Observe(durationSeconds)

	c.JSON(200, vacancy)
}

func (vc *VacancyController) DeleteVacancy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	err = vc.vacancyService.Delete(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete vacancy: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func (vc *VacancyController) EditVacancy(c *gin.Context) {
	var vacancy models.Vacancy

	err := c.ShouldBindJSON(&vacancy)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = vc.vacancyService.Update(&vacancy)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update vacancy: " + err.Error()})
		return
	}

	c.JSON(200, vacancy)
}
