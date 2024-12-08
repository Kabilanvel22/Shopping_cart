package main

import (
	"log"

	"example.com/go-backend/database"
	"example.com/go-backend/handlers"
	"example.com/go-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	
	database.ConnectDB()

	
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"}, 
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	
	public := r.Group("/")
	{
		public.POST("/users", handlers.CreateUser)     
		public.POST("/users/login", handlers.LoginUser) 
		public.GET("/items", handlers.ListItems)       
		public.POST("/items", handlers.CreateItem)       
	}

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware()) 
	{
		protected.GET("/users", handlers.ListUsers)    
		protected.POST("/carts", handlers.AddToCart)   
		protected.GET("/carts", handlers.ListCarts)    
		protected.POST("/orders", handlers.CreateOrder) 
		protected.GET("/orders", handlers.ListOrders)   
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Could not start the server:", err)
	}
}