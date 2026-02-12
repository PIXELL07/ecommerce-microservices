package main

import (
	"user-service/database"
	"user-service/handlers"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	//r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)

	r.Run(":8001")
}
