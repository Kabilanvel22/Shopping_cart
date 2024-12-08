package handlers

import (
	"example.com/go-backend/database"
	"example.com/go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ListCarts(c *gin.Context) {
	var carts []models.Cart
	userID, _ := c.Get("user_id")

	
	if err := database.DB.
		Where("user_id = ?", userID).
		Preload("User").             
		Preload("CartItems").        
		Preload("CartItems.Item").  
		Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve carts"})
		return
	}

	c.JSON(http.StatusOK, carts) 
}