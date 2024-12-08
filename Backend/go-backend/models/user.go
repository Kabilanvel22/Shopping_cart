package models


type User struct {
    ID       uint    `gorm:"primaryKey"`
    Username string  `gorm:"unique"`
    Password string
    Token    string
    Cart   Cart    `gorm:"constraint:OnDelete:CASCADE;"`
    Orders []Order `gorm:"foreignKey:UserID"`
}


