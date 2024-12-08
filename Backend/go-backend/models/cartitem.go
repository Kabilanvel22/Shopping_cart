package models


type CartItem struct {
    ID       uint  `gorm:"primaryKey"`
    CartID   uint  `gorm:"index"`
    Cart     Cart  `gorm:"constraint:OnDelete:CASCADE;"`
    ItemID   uint  `gorm:"index" json:"item_id"`
    Item     Item  `gorm:"constraint:OnDelete:CASCADE;"`
    Quantity int   `json:"quantity"` 
}
