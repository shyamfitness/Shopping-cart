package models

import "gorm.io/gorm"

// Order holds checkout info.
type Order struct {
	gorm.Model
	UserID uint
	Total  float64
}

// OrderItem stores the snapshot of cart items.
type OrderItem struct {
	gorm.Model
	OrderID  uint
	ItemID   uint
	Quantity int
	Price    float64
}

