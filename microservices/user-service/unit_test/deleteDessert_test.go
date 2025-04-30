package unit

import (
	"dessert-shop-backend/microservices/product-service/controllers"
	models2 "dessert-shop-backend/microservices/product-service/models"
	"dessert-shop-backend/microservices/user-service/database"
	"dessert-shop-backend/microservices/user-service/middleware"
	"dessert-shop-backend/microservices/user-service/models"
	"dessert-shop-backend/microservices/user-service/testhelper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteDessert(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	// 创建测试用户并生成 token
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := models.User{Username: "admin", Password: string(hashedPassword), Role: "admin"}
	database.DB.Create(&user)
	token := testhelper.GenerateTestToken(user)

	// 创建测试分类和甜点
	category := models2.Category{Name: "测试分类"}
	database.DB.Create(&category)
	dessert := models2.Dessert{Name: "测试蛋糕", Price: 12.5, CategoryID: category.ID}
	database.DB.Create(&dessert)

	// 设置路由
	router := gin.Default()
	router.DELETE("/api/desserts/:id", middleware.JWTAuthMiddleware(), controllers.DeleteDessert)

	req, _ := http.NewRequest("DELETE", "/api/desserts/"+formatID(dessert.ID), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "已删除")

	// 验证数据库中是否已删除
	var result models2.Dessert
	database.DB.First(&result, dessert.ID)
	assert.Equal(t, true, result.ID == 0)
}

// formatID 简单处理 uint -> string
func formatID(id uint) string {
	return fmt.Sprintf("%d", id)
}
