package controllers

import (
	"dessert-shop-backend/database"
	"dessert-shop-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	
	if err := database.DB.Select("ID", "name").Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取失败"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// 获取指定分类及其甜品
func GetCategoryWithDesserts(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	// 使用 Preload 加载关联甜品
	//.Select("id", "name", "price")
	result := database.DB.Preload("Desserts").First(&category, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.First(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "类别未找到"})
		return
	}
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "已删除"})
}
