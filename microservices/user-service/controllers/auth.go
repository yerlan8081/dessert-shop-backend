package controllers

import (
	"dessert-shop-backend/microservices/user-service/database"
	"dessert-shop-backend/microservices/user-service/models"
	"dessert-shop-backend/microservices/user-service/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效输入"})
		return
	}

	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库连接未初始化"})
		return
	}

	if user.Role == "" {
		user.Role = "user"
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}

	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	// 1. 绑定登录请求数据
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效输入"})
		return
	}

	// 2. 查找用户
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}

	// 3. 比较密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	// 4. 生成 JWT Token
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token 生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

	// 5. 使用 Resty 获取甜点数据
	//client := resty.New()
	//productResp, err := client.R().
	//	SetHeader("Content-Type", "application/json").
	//	SetHeader("Authorization", "Bearer "+token). // 添加这行
	//	Get("http://product-service:8082/api/desserts")
	//
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "获取甜点数据失败"})
	//	return
	//}

	// 6. 返回登录信息和甜点数据
	//c.JSON(http.StatusOK, gin.H{
	//	"token":    token,
	//	"user_id":  user.ID,
	//	"username": user.Username,
	//	"role":     user.Role,
	//	"products": productResp.String(), // 你可以根据需要解析为结构体
	//})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// 查找用户
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}

	// 绑定 JSON 到 user（覆盖原字段）
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 如果有密码，进行加密
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
			return
		}
		user.Password = string(hashedPassword)
	}

	// 保存更新
	database.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}
