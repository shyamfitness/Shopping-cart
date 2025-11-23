package main

import (
	"fmt"
	"log"

	"shopping-cart-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

// SetupDatabase opens the SQLite database and runs migrations.
func SetupDatabase() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}

	db, err := gorm.Open(sqlite.Open("shopping_cart.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	dbInstance = db
	fmt.Println("database ready")
	return dbInstance
}

