package controllers

import (
	"dessert-shop-backend/microservices/product-service/models"
	"dessert-shop-backend/microservices/user-service/database"
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

	c.JSON(http.StatusOK, gin.H{
		"message":  "分类创建成功",
		"category": category,
	})
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

func UpdateCategory(c *gin.Context) {
	id := c.Param("id") // 从 URL 中获取 ID 参数

	var category models.Category
	// 查询数据库中是否存在该分类
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该分类"})
		return
	}

	var updateData models.Category
	// 绑定更新数据
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新字段（这里只示例更新 Name 和 Description）
	category.Name = updateData.Name

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "分类更新成功",
		"category": category,
	})
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
