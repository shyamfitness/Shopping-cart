package models

import "gorm.io/gorm"

// Item represents something that can be bought.
type Item struct {
	gorm.Model
	Name  string
	Price float64
}

