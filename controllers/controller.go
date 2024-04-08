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
		return
	}
	//validation
	if err := models.ValidateVehicleData(&vehicle); err != nil { 
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error()})
		return
	} 
	database.DB.Create(&vehicle)   // in case there is no error do this
	c.JSON(http.StatusOK, vehicle)
}

func SearchVehiclePerID(c *gin.Context) {
	var vehicle models.Vehicle 
	id := c.Params.ByName("id")
	database.DB.First(&vehicle, id) //Find the first with this id, save on vehicle memory address 
	
	if vehicle.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Vehicle not found"})
			return
	}
	c.JSON(http.StatusOK, vehicle)  //show this vehicle on screen
}

func DeleteVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	id := c.Params.ByName("id")
	database.DB.Delete(&vehicle, id)
	c.JSON(http.StatusOK, gin.H {
		"data":"Vehicle deleted successfully"}) //delete this vehicle
}


func EditVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	id := c.Params.ByName("id")
	database.DB.First(&vehicle,id)

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error()})
		return	
	}

	//validation
	if err := models.ValidateVehicleData(&vehicle); err != nil { 
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error()})
		return
	} 

	database.DB.Model(&vehicle).UpdateColumns(vehicle)
	c.JSON(http.StatusOK, vehicle)
}

func SearchVehiclePerVin(c *gin.Context) {
	var vehicle models.Vehicle
	vin := c.Param("vin")
	database.DB.Where(models.Vehicle{Vin: vin}).First(&vehicle)

	if vehicle.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Vehicle not found"})
			return
	}

	c.JSON(http.StatusOK, vehicle)
}