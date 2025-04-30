package unit

import (
	"dessert-shop-backend/controllers"
	"dessert-shop-backend/database"
	"dessert-shop-backend/middleware"
	"dessert-shop-backend/models"
	"dessert-shop-backend/testhelper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteCartItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	// 创建用户和 token
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := models.User{Username: "user1", Password: string(hashedPassword), Role: "user"}
	database.DB.Create(&user)
	token := testhelper.GenerateTestToken(user)

	// 创建测试商品和购物车项
	category := models.Category{Name: "测试分类"}
	database.DB.Create(&category)

	dessert := models.Dessert{Name: "测试甜点", Price: 9.99, CategoryID: category.ID}
	database.DB.Create(&dessert)

	cartItem := models.CartItem{
		UserID:    user.ID,
		DessertID: dessert.ID,
		Quantity:  2,
	}
	database.DB.Create(&cartItem)

	// 设置路由
	router := gin.Default()
	router.DELETE("/api/cart/:id", middleware.JWTAuthMiddleware(), controllers.DeleteCartItem)

	req, _ := http.NewRequest("DELETE", "/api/cart/"+fmt.Sprint(cartItem.ID), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "已删除")

	// 验证是否已删除
	var result models.CartItem
	tx := database.DB.First(&result, cartItem.ID)
	assert.Error(t, tx.Error) // 应该找不到记录
}
