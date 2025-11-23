package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shopping-cart-backend/controllers"
)

// SetupRouter wires routes to controllers.
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	userController := controllers.NewUserController(db)
	itemController := controllers.NewItemController(db)
	cartController := controllers.NewCartController(db)
	orderController := controllers.NewOrderController(db)

	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.ListUsers)
	router.POST("/users/login", userController.Login)

	router.POST("/items", itemController.CreateItem)
	router.GET("/items", itemController.ListItems)

	router.GET("/carts", cartController.ListCarts)
	router.GET("/orders", orderController.ListOrders)

	auth := router.Group("/")
	auth.Use(TokenMiddleware(db))
	auth.POST("/carts", cartController.AddToCart)
	auth.POST("/orders", orderController.CreateOrder)

	return router
}

