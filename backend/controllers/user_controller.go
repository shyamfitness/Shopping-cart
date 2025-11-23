package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shopping-cart-backend/models"
)

// UserController keeps a reference to the DB.
type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

type userInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var body userInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	user := models.User{
		Username: body.Username,
		Password: body.Password,
	}

	if err := uc.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) ListUsers(c *gin.Context) {
	var users []models.User
	uc.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) Login(c *gin.Context) {
	var body userInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	var user models.User
	if err := uc.DB.Where("username = ?", body.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if user.Password != body.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	token := fmt.Sprintf("token-%d", time.Now().UnixNano())
	user.CurrentToken = token
	uc.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

