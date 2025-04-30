package unit_test

import (
	"bytes"
	"dessert-shop-backend/controllers"
	"dessert-shop-backend/database"
	"dessert-shop-backend/models"
	"dessert-shop-backend/testhelper"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	testhelper.SetupTestMain(m)
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// 设置测试数据库
	database.DB = database.TestDB

	router := gin.Default()
	router.POST("/api/register", controllers.Register)

	user := models.User{
		Username: "testuser",
		Password: "testpass",
		Role:     "",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "注册成功")

	testhelper.CleanDatabase(database.DB)
}
