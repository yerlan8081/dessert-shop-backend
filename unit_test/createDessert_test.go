package unit

import (
	"bytes"
	"dessert-shop-backend/controllers"
	"dessert-shop-backend/database"
	"dessert-shop-backend/middleware"
	"dessert-shop-backend/models"
	"dessert-shop-backend/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDessert(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := models.User{Username: "admin", Password: string(hashedPassword), Role: "admin"}
	database.DB.Create(&user)

	// 创建分类
	category := models.Category{Name: "招牌"}
	database.DB.Create(&category)

	token := testhelper.GenerateTestToken(user)

	router := gin.Default()
	router.POST("/api/desserts", middleware.JWTAuthMiddleware(), controllers.CreateDessert)

	payload := map[string]interface{}{
		"name":        "草莓蛋糕",
		"price":       25.5,
		"description": "好吃的草莓蛋糕",
		"category_id": category.ID,
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/desserts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "甜品创建成功")
}
