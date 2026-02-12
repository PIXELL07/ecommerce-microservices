package handlers

import (
	"net/http"

	"user-service/models"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all users

func GetUsers(c *gin.Context) {

	var users []models.User

	// Logic to fetch users from the database

	c.JSON(http.StatusOK, users)

}
