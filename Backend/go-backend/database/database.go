package database

import (
	"log"

	"example.com/go-backend/models"
	"gorm.io/driver/mysql" 
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB() {
	var err error
	
	dsn := "root:tiger#123@tcp(127.0.0.1:3306)/shopping?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	
	err = DB.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Cart{},
		&models.Item{},
		&models.CartItem{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")
}
