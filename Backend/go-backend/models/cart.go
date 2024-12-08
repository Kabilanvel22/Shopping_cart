package models

import "time"

type Cart struct {
    ID        uint       `gorm:"primaryKey"`
    UserID    uint       `gorm:"index;not null"`  
    User      *User      `gorm:"constraint:OnDelete:CASCADE;"`  
    CartItems []CartItem `gorm:"foreignKey:CartID"`  
    Order     *Order     `gorm:"constraint:OnDelete:SET NULL;"` 
    CreatedAt time.Time
}
