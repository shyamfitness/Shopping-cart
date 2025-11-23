package main

import (
	"log"

	"shopping-cart-backend/routes"
)

func main() {
	db := SetupDatabase()
	router := routes.SetupRouter(db)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

