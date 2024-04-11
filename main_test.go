package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/tiagoc0sta/api-go-gin/controllers"
)

func SetupOfTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

// Test created to fail
func TestVerifyStatusCodeOfGreetingWithParameter(t *testing.T) {
	r := SetupOfTestRoutes() // Call the function to get the gin.Engine instance
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/Tiago", nil)
	response := httptest.NewRecorder() // NewRecorder is an Interface of Response Writer
	r.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Fatalf("Status error: received value was %d and the expected was %d", response.Code, http.StatusOK)
	}
}

