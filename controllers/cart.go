package controllers

import (
	"dessert-shop-backend/database"
	"dessert-shop-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加商品到购物车
func AddToCart(c *gin.Context) {
	var input struct {
		DessertID uint `json:"dessert_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 获取用户ID（假设从 JWT 中间件中读取）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	cartItem := models.CartItem{
		UserID:    userID.(uint),
		DessertID: input.DessertID,
		Quantity:  input.Quantity,
	}

	if err := database.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加失败"})
		return
	}

	database.DB.Preload("Dessert").First(&cartItem, cartItem.ID)

	c.JSON(http.StatusOK, gin.H{"message": "已添加到购物车", "item": cartItem})
}

// 获取用户购物车内容
func GetCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	var cartItems []models.CartItem
	err := database.DB.Preload("Dessert").Where("user_id = ?", userID).Find(&cartItems).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cartItems})
}

// 从购物车中删除一项
func DeleteCartItem(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	id := c.Param("id")

	var item models.CartItem
	result := database.DB.First(&item, id)
	if result.Error != nil || item.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{"error": "购物车项不存在或无权限"})
		return
	}

	database.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "已删除"})
}
