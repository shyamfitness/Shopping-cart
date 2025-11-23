package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shopping-cart-backend/models"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	rawUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user"})
		return
	}
	user := rawUser.(models.User)

	var cart models.Cart
	if err := oc.DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	var items []models.CartItem
	oc.DB.Where("cart_id = ?", cart.ID).Find(&items)
	if len(items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart is empty"})
		return
	}

	order := models.Order{UserID: user.ID}
	var total float64
	oc.DB.Create(&order)

	for _, cartItem := range items {
		var storeItem models.Item
		if err := oc.DB.First(&storeItem, cartItem.ItemID).Error; err != nil {
			continue
		}

		total += float64(cartItem.Quantity) * storeItem.Price
		orderItem := models.OrderItem{
			OrderID:  order.ID,
			ItemID:   storeItem.ID,
			Quantity: cartItem.Quantity,
			Price:    storeItem.Price,
		}
		oc.DB.Create(&orderItem)
	}

	order.Total = total
	oc.DB.Save(&order)

	oc.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})

	c.JSON(http.StatusOK, gin.H{"message": "order created", "order_id": order.ID})
}

func (oc *OrderController) ListOrders(c *gin.Context) {
	var orders []models.Order
	oc.DB.Find(&orders)

	var orderItems []models.OrderItem
	oc.DB.Find(&orderItems)

	c.JSON(http.StatusOK, gin.H{
		"orders":     orders,
		"orderItems": orderItems,
	})
}

