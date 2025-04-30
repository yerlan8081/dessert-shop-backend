package unit_test

import (
	"bytes"
	"dessert-shop-backend/microservices/user-service/controllers"
	"dessert-shop-backend/microservices/user-service/database"
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

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	// 创建测试用户
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
	database.DB.Create(&models.User{
		Username: "testuser",
		Password: string(hashedPassword),
	})

	router := gin.Default()
	router.POST("/api/login", controllers.Login)

	credentials := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	jsonValue, _ := json.Marshal(credentials)
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println("响应体:", w.Body.String())
	assert.Contains(t, w.Body.String(), "token")
}
