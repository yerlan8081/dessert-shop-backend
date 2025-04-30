package controllers

import (
	"dessert-shop-backend/microservices/product-service/models"
	"dessert-shop-backend/microservices/user-service/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建甜品
func CreateDessert(c *gin.Context) {
	var dessert models.Dessert
	if err := c.ShouldBindJSON(&dessert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&dessert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"dessert": dessert,
		"message": "甜品创建成功",
	})
}

// 获取所有甜品
func GetDesserts(c *gin.Context) {
	var desserts []models.Dessert

	// 分页参数
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	nameFilter := c.Query("name")        // 模糊搜索
	categoryID := c.Query("category_id") // 分类筛选

	var pageInt, limitInt int
	fmt.Sscan(page, &pageInt)
	fmt.Sscan(limit, &limitInt)

	offset := (pageInt - 1) * limitInt

	// 初始化查询
	query := database.DB.Model(&models.Dessert{})

	if nameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+nameFilter+"%")
	}

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// 执行查询
	query.Limit(limitInt).Offset(offset).Find(&desserts)

	c.JSON(http.StatusOK, gin.H{
		"page":     pageInt,
		"limit":    limitInt,
		"desserts": desserts,
		"message":  "甜点",
	})
}

// 根据 ID 获取单个甜品
func GetDessert(c *gin.Context) {
	var dessert models.Dessert
	id := c.Param("id")
	if err := database.DB.First(&dessert, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "甜品未找到"})
		return
	}
	c.JSON(http.StatusOK, dessert)
}
func GetDessertsByCategory(c *gin.Context) {
	categoryID := c.Param("id")

	var desserts []models.Dessert

	result := database.DB.Where("category_id = ?", categoryID).Find(&desserts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, desserts)
}

// 更新甜品
func UpdateDessert(c *gin.Context) {
	var dessert models.Dessert
	id := c.Param("id")

	if err := database.DB.First(&dessert, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "甜品未找到"})
		return
	}

	if err := c.ShouldBindJSON(&dessert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&dessert)
	c.JSON(http.StatusOK, dessert)
}

// 删除甜品
func DeleteDessert(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Dessert{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "已删除"})
}
