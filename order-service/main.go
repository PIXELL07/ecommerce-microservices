package main

import (
	"order-service/database"
	"order-service/handlers"
	"order-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Order{})

	r := gin.Default()
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)

	r.Run(":8003")
}
