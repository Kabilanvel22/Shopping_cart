package models

import "time"

type Order struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `gorm:"index"; json:"user_id"`
    User      *User     `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
    CartID    *uint     `gorm:"uniqueIndex;null" json:"cart_id"` 
    Cart      *Cart     `gorm:"constraint:OnDelete:SET NULL;" json:"cart"`
    CreatedAt time.Time `json:"created_at"`
}
