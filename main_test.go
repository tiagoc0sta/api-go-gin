package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tiagoc0sta/api-go-gin/controllers"
	"github.com/tiagoc0sta/api-go-gin/database"
	"github.com/tiagoc0sta/api-go-gin/models"
)

var ID int

func SetupOfTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateVehicleMock(){
	vehicle:= models.Vehicle{Name:  "Vehicle name test",	Vin: "3N1CB51D65L461013",  Brand: "Vehicle brand test"}
	database.DB.Create(&vehicle)
	ID = int(vehicle.ID)
}

func DeleteVehicleMock(){
	var vehicle models.Vehicle
	database.DB.Delete(&vehicle, ID)

}

// Test created to fail
func TestVerifyStatusCodeOfGreetingWithParameter(t *testing.T) {
	r := SetupOfTestRoutes() // Call the function to get the gin.Engine instance
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/Tiago", nil)
	response := httptest.NewRecorder() // NewRecorder is an Interface of Response Writer
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
	mockOfResponse := `{"API says:":"What's up Tiago, Is everything alright?"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, mockOfResponse, string(responseBody))
	fmt.Println(string(responseBody))
	fmt.Println(mockOfResponse)
}

func TestListAllVehiclesHandler(t *testing.T){
	database.ConnectToDatabase()
	CreateVehicleMock()
	defer DeleteVehicleMock()
	r := SetupOfTestRoutes()
	r.GET("/vehicles", controllers.ShowAllVehicles)
	req, _ := http.NewRequest("GET","/vehicles", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	//fmt.Println(response.Body)
}

/*func TestSearchVehiclePerVinHandler(t *testing.T){
	database.ConnectToDatabase()
	CreateVehicleMock()
	defer DeleteVehicleMock()
	r := SetupOfTestRoutes()
	r.GET("vehicles/vin/:vin", controllers.SearchVehiclePerVin)
	req, _ := http.NewRequest("GET","vehicles/vin/3N1CB51D65L461013", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}*/

//diferent kind of test. Verifying if the vehicle has the mock information on it
func TestSearchVehiclePerIDHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateVehicleMock()
	defer DeleteVehicleMock()
	r := SetupOfTestRoutes()
	r.GET("/vehicles/:id", controllers.SearchVehiclePerID)
	pathOfSearch := "/vehicles/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathOfSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var vehicleMock models.Vehicle
	json.Unmarshal(response.Body.Bytes(), &vehicleMock)
	assert.Equal(t, "Vehicle name test", vehicleMock.Name)
	assert.Equal(t, "3N1CB51D65L461013", vehicleMock.Vin)
	assert.Equal(t, "Vehicle brand test", vehicleMock.Brand)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteVehicleHandler(t *testing.T){
	database.ConnectToDatabase()
	CreateVehicleMock()
	r := SetupOfTestRoutes()
	r.DELETE("/vehicles/:id", controllers.DeleteVehicle)
	pathOfSearch := "/vehicles/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathOfSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditVehicleHandler(t *testing.T){
	database.ConnectToDatabase()
	CreateVehicleMock()
	defer DeleteVehicleMock()
	r := SetupOfTestRoutes()
	r.PATCH("/vehicles/:id", controllers.EditVehicle)
	vehicle := models.Vehicle{Name:  "Vehicle name test",	Vin: "181CB51D65L461013",  Brand: "Vehicle brand test"}
	valueJson, _ := json.Marshal(vehicle)
	pathToEdit := "/vehicles/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathToEdit, bytes.NewBuffer(valueJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var vehicleMockUpdated models.Vehicle
	json.Unmarshal(response.Body.Bytes(), &vehicleMockUpdated)
	assert.Equal(t, "Vehicle name test", vehicleMockUpdated.Name)
	assert.Equal(t, "181CB51D65L461013", vehicleMockUpdated.Vin)
	assert.Equal(t, "Vehicle brand test", vehicleMockUpdated.Brand)

}