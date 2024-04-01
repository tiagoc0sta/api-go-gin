package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoc0sta/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.Run()
}