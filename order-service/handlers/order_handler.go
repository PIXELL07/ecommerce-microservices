// CreateOrder handles the creation of a new order.
// It binds the incoming JSON request to an Order model,
// validates the user and product IDs, and if valid,
// saves the order to the database. If the user or product
// is invalid, it responds with a 400 Bad Request error.
// On successful creation, it returns a 201 Created status
// along with the created order object.

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
