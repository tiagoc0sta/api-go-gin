package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoc0sta/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/vehicles", controllers.ShowAllVehicles)
	r.GET("/:name", controllers.Greeting)
	r.POST("/vehicles", controllers.CreateNewVehicle)
	r.GET("/vehicles/:id", controllers.SearchVehiclePerID)
	r.DELETE("/vehicles/:id", controllers.DeleteVehicle)
	r.PATCH("/vehicles/:id", controllers.EditVehicle)
	r.GET("vehicles/vin/:vin", controllers.SearchVehiclePerVin)
	r.Run()
}