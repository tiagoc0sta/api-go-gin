package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"github.com/tiagoc0sta/api-go-gin/database"
	"github.com/tiagoc0sta/api-go-gin/models"
)

// ShowAllVehicles godoc
//
//	@Summary		show all vehicles
//	@Description	route to show all products
//	@Tags			vehicles
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Vehicle
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/vehicles [get]
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

// CreateNewVehicle godoc
//
//	@Summary		Add new vehicle
//	@Description	route to create a new products
//	@Tags			vehicles
//	@Accept			json
//	@Produce		json
//	@Param			vehicle body models.Vehicle true	"Vehicle model" 
//	@Success		200	{object}	models.Vehicle
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/vehicles [post]
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


// SearchVehiclePerID godoc
//
//	@Summary		Show vehicle per ID
//	@Description	Route to show vehicle per ID
//	@Tags			vehicles
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Vehicle ID"
//	@Success		200	{object}	models.Vehicle
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/vehicles/{id} [get]
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


// DeleteVehicle godoc
//
//	@Summary		Delete vehicle per ID
//	@Description	Route to delete vehicle per ID
//	@Tags			vehicles
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Vehicle ID"
//	@Success		200	{object}	models.Vehicle
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/vehicles/{id} [delete]
func DeleteVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	id := c.Params.ByName("id")
	database.DB.Delete(&vehicle, id)
	c.JSON(http.StatusOK, gin.H {
		"data":"Vehicle deleted successfully"}) //delete this vehicle
}


// EditVehicle godoc
//
//	@Summary		Update vehicle per ID
//	@Description	Route to update vehicle per ID
//	@Tags			vehicles
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Vehicle ID"
//  @Param vehicle body models.Vehicle true "vehiccle update"
//	@Success		200	{object}	models.Vehicle
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/vehicles/{id} [patch]
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

//show frontend - index.html
func ShowIndexPage(c *gin.Context) {
	var vehicles []models.Vehicle //create a variable with a slice of vehicles 
	database.DB.Find(&vehicles)  //database will find all vehicles and save into vehicles variable(slice)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"vehicles":vehicles,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}