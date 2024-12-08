package handlers

import (
	"log"
	"net/http"

	"example.com/go-backend/database"
	"example.com/go-backend/models"

	"github.com/gin-gonic/gin"
)


func AddToCart(c *gin.Context) {
	var cartItem models.CartItem
	userID, exists := c.Get("user_id") 

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		log.Printf("Error binding request data: %v", err) 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	
	log.Printf("Received item_id: %d, Quantity: %d", cartItem.ItemID, cartItem.Quantity)

	
	if cartItem.ItemID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item_id"})
		return
	}

	
	var cart models.Cart
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		
		cart = models.Cart{
			UserID: userID.(uint), 
		}
		if err := database.DB.Create(&cart).Error; err != nil {
			log.Printf("Failed to create cart for user ID %d: %v", userID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
			return
		}
		log.Printf("Created new cart with ID %d for user ID %d", cart.ID, userID)
	}

	
	var item models.Item
	if err := database.DB.First(&item, cartItem.ItemID).Error; err != nil {
		log.Printf("Error finding item with ID %d: %v", cartItem.ItemID, err) 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item does not exist"})
		return
	}


	var existingCartItem models.CartItem
	if err := database.DB.Where("cart_id = ? AND item_id = ?", cart.ID, cartItem.ItemID).First(&existingCartItem).Error; err == nil {
		
		existingCartItem.Quantity += cartItem.Quantity
		if err := database.DB.Save(&existingCartItem).Error; err != nil {
			log.Printf("Failed to update quantity for item ID %d in cart ID %d: %v", cartItem.ItemID, cart.ID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item quantity"})
			return
		}
		log.Printf("Updated quantity for item ID %d in cart ID %d", cartItem.ItemID, cart.ID)
		c.JSON(http.StatusOK, gin.H{"message": "Item quantity updated", "cart_id": cart.ID}) 
		return
	}

	cartItem.CartID = cart.ID 
	if err := database.DB.Create(&cartItem).Error; err != nil {
		log.Printf("Failed to add item ID %d to cart ID %d: %v", cartItem.ItemID, cart.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}
	log.Printf("Added item ID %d to cart ID %d", cartItem.ItemID, cart.ID)

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart", "cart_id": cart.ID})
}
