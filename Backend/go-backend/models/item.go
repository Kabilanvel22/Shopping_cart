package models


type Item struct {
    ID    uint    `gorm:"primaryKey"`
    Name  string
    Price float64
    CartItems []CartItem `gorm:"foreignKey:ItemID"`
}
