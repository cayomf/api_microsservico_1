package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/models"
)

// updateEntrega responds with the list of all vacancies as JSON.
// updateEntrega             godoc
// @Summary      Atualiza a nota na entrega do aluno
// @Description  Atualiza a nota na entrega do aluno
// @Tags         nota
// @Produce      json
// @Param        body     body     models.EntregaRequest     true        "EntregaRequest"
// @Success      200  {array}  models.EntregaRequest
// @Router       /updateEntrega [post]
func ShowAllVacancies(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Vacancy
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find vacancy by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowLastVacancies(c *gin.Context) {
	quantity, _ := strconv.ParseInt(c.Param("quantity"), 10, 64)
	db := database.GetDatabase()
	var p []models.Vacancy
	err := db.Limit(int(quantity)).Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot get last elements: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowVacancy(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.Vacancy
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find vacancy by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateVacancy(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Vacancy

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create vacancy: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteVacancy(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Vacancy{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete vacancy: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EditVacancy(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Vacancy

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create vacancy: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
