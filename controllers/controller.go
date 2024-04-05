package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiagoc0sta/api-go-gin/database"
	"github.com/tiagoc0sta/api-go-gin/models"
)

func ShowAllVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	database.DB.Find(&vehicles)
	c.JSON(200, vehicles)
}

func Greeting (c *gin.Context) {
	nome := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:" : "What's up " + nome + ", Is everything alright?",
	}) 	
}

func CreateNewVehicle(c *gin.Context) {
	var vehicle models.Vehicle

	if err := c.ShouldBindJSON(&vehicle); err != nil { //in case there is error (error not empty)
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error()})
	}
	database.DB.Create(&vehicle)   // in case there is no error do this
	c.JSON(http.StatusOK, vehicle)

}