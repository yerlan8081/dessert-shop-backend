package unit

import (
	"bytes"
	"dessert-shop-backend/microservices/product-service/controllers"
	models2 "dessert-shop-backend/microservices/product-service/models"
	controllers2 "dessert-shop-backend/microservices/user-service/controllers"
	"dessert-shop-backend/microservices/user-service/database"
	"dessert-shop-backend/microservices/user-service/middleware"
	"dessert-shop-backend/microservices/user-service/models"
	"dessert-shop-backend/microservices/user-service/testhelper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddToCart(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	// 1. 创建用户并加密密码
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
	user := models.User{Username: "testuser", Password: string(hashedPassword)}
	database.DB.Create(&user)

	// 2. 登录获取 token
	router := gin.Default()
	router.POST("/api/login", controllers2.Login)
	loginPayload := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	body, _ := json.Marshal(loginPayload)
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	fmt.Println("Login response:", resp)

	token, ok := resp["token"].(string)
	assert.True(t, ok, "token 字段类型错误或不存在")
	assert.NotEmpty(t, token, "token 不应为空")

	// 3. 创建 Category
	category := models2.Category{Name: "Cakes"}
	database.DB.Create(&category)

	// 4. 创建 Dessert，使用上面 category 的 ID
	dessert := models2.Dessert{
		Name:       "cake",
		Price:      10,
		CategoryID: category.ID,
	}
	database.DB.Create(&dessert)

	// 5. 添加商品到购物车
	authorized := gin.Default()
	cartGroup := authorized.Group("/api/cart")
	cartGroup.Use(middleware.JWTAuthMiddleware())
	cartGroup.POST("", controllers.AddToCart)

	cartPayload := map[string]interface{}{
		"dessert_id": dessert.ID,
		"quantity":   2,
	}
	cartBody, _ := json.Marshal(cartPayload)
	cartReq, _ := http.NewRequest("POST", "/api/cart", bytes.NewBuffer(cartBody))
	cartReq.Header.Set("Content-Type", "application/json")
	cartReq.Header.Set("Authorization", "Bearer "+token)

	cartResp := httptest.NewRecorder()
	authorized.ServeHTTP(cartResp, cartReq)

	assert.Equal(t, 200, cartResp.Code)
	assert.Contains(t, cartResp.Body.String(), "已添加到购物车")
}
