package main

import (
	"github.com/tiagoc0sta/api-go-gin/database"

	"github.com/tiagoc0sta/api-go-gin/routes"
)

func main() {
	database.ConnectToDatabase() //This comand connects to db postgres
	routes.HandleRequests()
}