package models

import "gorm.io/gorm"

// User holds login info and the active token.
type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Password     string
	CurrentToken string
}

