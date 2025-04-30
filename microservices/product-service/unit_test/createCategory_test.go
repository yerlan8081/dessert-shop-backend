package unit

import (
	"bytes"
	"dessert-shop-backend/microservices/product-service/controllers"
	"dessert-shop-backend/microservices/user-service/database"
	"dessert-shop-backend/microservices/user-service/middleware"
	"dessert-shop-backend/microservices/user-service/models"
	"dessert-shop-backend/microservices/user-service/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := models.User{Username: "admin", Password: string(hashedPassword), Role: "admin"}
	database.DB.Create(&user)

	token := testhelper.GenerateTestToken(user)

	router := gin.Default()
	router.POST("/api/categories", middleware.JWTAuthMiddleware(), controllers.CreateCategory)

	payload := map[string]interface{}{
		"name": "新品推荐",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "分类创建成功")
}
