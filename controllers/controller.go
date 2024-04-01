package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoc0sta/api-go-gin/models"
)

func ShowAllVehicles(c *gin.Context) {
	c.JSON(200,models.Vehicles)
}

func Greeting (c *gin.Context) {
	nome := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:" : "What's up " + nome + ", Is everything alright?",
	}) 	
}