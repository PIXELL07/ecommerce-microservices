package main

import (
	// "os" import removed coz not used
	"product-service/database"
	"product-service/handlers"
	"product-service/models"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	database.DB.AutoMigrate(&models.Product{})

	r := gin.Default()

	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProductByID)

	r.Run(":8002")
}
