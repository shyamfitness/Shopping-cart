package models

import "gorm.io/gorm"

// Cart is tied to a single user.
type Cart struct {
	gorm.Model
	UserID uint
}

// CartItem links carts and items.
type CartItem struct {
	gorm.Model
	CartID   uint
	ItemID   uint
	Quantity int
}

