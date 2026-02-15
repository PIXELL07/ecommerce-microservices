package handlers

import (
	"net/http"
	"order-service/database"
	"order-service/models"
	"order-service/services"

	"github.com/gin-gonic/gin"
)

// CreateOrder handles the creation of a new order.
// It validates the user and product before saving the order to the database.

func CreateOrder(c *gin.Context) {
	var order models.Order
	c.BindJSON(&order)

	if !services.ValidateUser(order.UserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		return
	}

	if !services.ValidateProduct(order.ProductID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product"})
		return
	}

	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

// GetOrders handles fetching all orders

func GetOrders(c *gin.Context) {

	var orders []models.Order

	// Implementation for fetching orders

	c.JSON(http.StatusOK, orders)

}
