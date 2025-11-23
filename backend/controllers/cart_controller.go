package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shopping-cart-backend/models"
)

type CartController struct {
	DB *gorm.DB
}

func NewCartController(db *gorm.DB) *CartController {
	return &CartController{DB: db}
}

type cartInput struct {
	ItemID   uint `json:"item_id"`
	Quantity int  `json:"quantity"`
}

func (cc *CartController) AddToCart(c *gin.Context) {
	rawUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user"})
		return
	}
	user := rawUser.(models.User)

	var body cartInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if body.Quantity <= 0 {
		body.Quantity = 1
	}

	var cart models.Cart
	err := cc.DB.Where("user_id = ?", user.ID).First(&cart).Error
	if err != nil {
		cart = models.Cart{UserID: user.ID}
		cc.DB.Create(&cart)
	}

	var item models.Item
	if err := cc.DB.First(&item, body.ItemID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
		return
	}

	var cartItem models.CartItem
	if err := cc.DB.Where("cart_id = ? AND item_id = ?", cart.ID, item.ID).First(&cartItem).Error; err != nil {
		cartItem = models.CartItem{
			CartID:   cart.ID,
			ItemID:   item.ID,
			Quantity: body.Quantity,
		}
		cc.DB.Create(&cartItem)
	} else {
		cartItem.Quantity += body.Quantity
		cc.DB.Save(&cartItem)
	}

	c.JSON(http.StatusOK, gin.H{"message": "item added"})
}

func (cc *CartController) ListCarts(c *gin.Context) {
	var carts []models.Cart
	cc.DB.Find(&carts)

	var cartItems []models.CartItem
	cc.DB.Find(&cartItems)

	c.JSON(http.StatusOK, gin.H{
		"carts":     carts,
		"cartItems": cartItems,
	})
}

