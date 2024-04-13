package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoc0sta/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*") //inform gin to load index.html inside templates
	r.Static("/assets", "./assets") //inform gin to load Static files (css)
	r.GET("/:name", controllers.Greeting)
	r.GET("/vehicles", controllers.ShowAllVehicles)	
	r.POST("/vehicles", controllers.CreateNewVehicle)
	r.GET("/vehicles/:id", controllers.SearchVehiclePerID)
	r.DELETE("/vehicles/:id", controllers.DeleteVehicle)
	r.PATCH("/vehicles/:id", controllers.EditVehicle)
	r.GET("vehicles/vin/:vin", controllers.SearchVehiclePerVin)
	r.GET("/index", controllers.ShowIndexPage)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}

