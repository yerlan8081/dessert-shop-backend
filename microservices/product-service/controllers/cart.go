package controllers

import (
	"dessert-shop-backend/microservices/product-service/database"
	"dessert-shop-backend/microservices/product-service/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func GetCartItemByID(c *gin.Context) {
	// 获取 URL 中的 cart item ID
	id := c.Param("id")

	var cartItem models.CartItem
	err := database.DB.Preload("Dessert").First(&cartItem, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "购物车项不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart_item": cartItem})
}

func UpdateCartItem(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Quantity int `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 获取用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	var cartItem models.CartItem
	err := database.DB.First(&cartItem, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "购物车项不存在"})
		return
	}

	// 确保该购物车项属于当前用户
	if cartItem.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限修改该购物车项"})
		return
	}

	// 更新数量
	cartItem.Quantity = input.Quantity
	if err := database.DB.Save(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	// 加载关联的 Dessert 信息
	database.DB.Preload("Dessert").First(&cartItem, cartItem.ID)

	c.JSON(http.StatusOK, gin.H{"message": "购物车项已更新", "item": cartItem})
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
