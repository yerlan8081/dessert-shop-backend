package unit

import (
	"dessert-shop-backend/controllers"
	"dessert-shop-backend/database"
	"dessert-shop-backend/models"
	"dessert-shop-backend/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCategories(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = database.TestDB
	testhelper.CleanDatabase(database.DB)

	// 创建一些测试分类
	category1 := models.Category{Name: "新品推荐"}
	category2 := models.Category{Name: "热销甜品"}
	database.DB.Create(&category1)
	database.DB.Create(&category2)

	// 设置路由
	router := gin.Default()
	router.GET("/api/categories", controllers.GetCategories)

	// 发起请求
	req, _ := http.NewRequest("GET", "/api/categories", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 解析响应
	var resp []models.Category
	err := json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code)
	assert.Len(t, resp, 2)
	assert.Equal(t, "新品推荐", resp[0].Name)
	assert.Equal(t, "热销甜品", resp[1].Name)
}
