package unit

import (
	"dessert-shop-backend/controllers"
	"dessert-shop-backend/database"
	"dessert-shop-backend/middleware"
	"dessert-shop-backend/models"
	"dessert-shop-backend/testhelper"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDesserts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	// 创建测试用户
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := models.User{Username: "admin", Password: string(hashedPassword), Role: "admin"}
	database.DB.Create(&user)
	token := testhelper.GenerateTestToken(user)

	// 创建分类（因为desserts需要category_id）
	category := models.Category{Name: "推荐"}
	database.DB.Create(&category)

	// 插入甜点数据（有category_id，但不含图片字段）
	dessert1 := models.Dessert{Name: "巧克力蛋糕", Price: 19.99, Description: "浓郁巧克力", CategoryID: category.ID}
	dessert2 := models.Dessert{Name: "草莓奶酪", Price: 15.50, Description: "清新草莓", CategoryID: category.ID}
	database.DB.Create(&dessert1)
	database.DB.Create(&dessert2)

	// 设置路由
	router := gin.Default()
	router.GET("/api/desserts", middleware.JWTAuthMiddleware(), controllers.GetDesserts)

	// 发起请求
	req, _ := http.NewRequest("GET", "/api/desserts?page=1&limit=10", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 响应结构体（不含图片）
	type DessertResponse struct {
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
	}

	type ResponseBody struct {
		Page     int               `json:"page"`
		Limit    int               `json:"limit"`
		Desserts []DessertResponse `json:"desserts"`
		Message  string            `json:"message"`
	}

	var resp ResponseBody
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Nil(t, err)

	// 验证响应
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "甜点", resp.Message)
	assert.Equal(t, 1, resp.Page)
	assert.Equal(t, 10, resp.Limit)
	assert.Len(t, resp.Desserts, 2)

	assert.Equal(t, "巧克力蛋糕", resp.Desserts[0].Name)
	assert.Equal(t, 19.99, resp.Desserts[0].Price)

	assert.Equal(t, "草莓奶酪", resp.Desserts[1].Name)
	assert.Equal(t, 15.50, resp.Desserts[1].Price)
}
