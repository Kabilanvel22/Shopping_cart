package handlers

import (
	"log"
	"net/http"

	"example.com/go-backend/database"
	"example.com/go-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

  
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    log.Printf("Received CartID: %d", order.CartID)

   
    var cart models.Cart
    if err := database.DB.Where("id = ? AND user_id = ?", order.CartID, userID).First(&cart).Error; err != nil {
        log.Printf("Error finding cart with ID %d for user ID %d: %v", order.CartID, userID, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart does not exist or does not belong to the user"})
        return
    }

    order.UserID = userID.(uint) 
  
    if err := database.DB.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    if err := database.DB.Preload("User").Preload("Cart").Preload("Items").First(&order, order.ID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to preload order details"})
        return
    }

   
    c.JSON(http.StatusOK, order)
}




func ListOrders(c *gin.Context) {
    var orders []models.Order
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    uid, ok := userID.(uint)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := database.DB.Where("user_id = ?", uid).Preload("Cart").Preload("Cart.CartItems").Preload("Cart.CartItems.Item").Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
        return
    }

    c.JSON(http.StatusOK, orders)
}


