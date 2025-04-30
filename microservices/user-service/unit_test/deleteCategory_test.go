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

func TestDeleteCategory(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	// 创建管理员用户和 token
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := models.User{Username: "admin", Password: string(hashedPassword), Role: "admin"}
	database.DB.Create(&user)
	token := testhelper.GenerateTestToken(user)

	// 创建测试分类
	category := models2.Category{Name: "测试分类"}
	database.DB.Create(&category)

	// 设置路由
	router := gin.Default()
	router.DELETE("/api/categories/:id", middleware.JWTAuthMiddleware(), controllers.DeleteCategory)

	req, _ := http.NewRequest("DELETE", "/api/categories/"+fmt.Sprint(category.ID), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "已删除")

	// 检查数据库中该分类是否已被删除
	var result models2.Category
	tx := database.DB.First(&result, category.ID)
	assert.Error(t, tx.Error) // 记录应不存在
}
