package main

import (
	"github.com/tiagoc0sta/api-go-gin/database"
	"github.com/tiagoc0sta/api-go-gin/models"
	"github.com/tiagoc0sta/api-go-gin/routes"
)

func main() {
	database.ConnectToDatabase() //This comand connects to db postgres
	models.Vehicles = []models.Vehicle{
		{Name: "Corolla", Vin:"213456789", Brand:"Toyota"},
		{Name: "Civic", Vin:"159437862", Brand:"Honda"},
	}

	routes.HandleRequests()
}