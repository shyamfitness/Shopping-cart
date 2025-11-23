package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shopping-cart-backend/models"
)

type ItemController struct {
	DB *gorm.DB
}

func NewItemController(db *gorm.DB) *ItemController {
	return &ItemController{DB: db}
}

type itemInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (ic *ItemController) CreateItem(c *gin.Context) {
	var body itemInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	item := models.Item{Name: body.Name, Price: body.Price}
	if err := ic.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (ic *ItemController) ListItems(c *gin.Context) {
	var items []models.Item
	ic.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

