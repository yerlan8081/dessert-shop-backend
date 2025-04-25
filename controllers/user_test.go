package controllers

//
//import (
//	"bytes"
//	"dessert-shop-backend/models"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//	"log"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//var DB *gorm.DB
//
//// 用于初始化数据库连接
//func setupTestDB() {
//	var err error
//	// 使用 SQLite 内存数据库进行测试
//	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("无法连接到数据库: %v", err)
//	}
//
//	// 自动迁移模型
//	err = DB.AutoMigrate(&models.User{})
//	if err != nil {
//		log.Fatalf("无法自动迁移模型: %v", err)
//	}
//}
//
//// 设置路由
//func setupRouter() *gin.Engine {
//	r := gin.Default()
//	r.POST("/register", Register)
//	return r
//}
//
//func TestRegister(t *testing.T) {
//	// 初始化数据库
//	setupTestDB()
//
//	// 创建一个路由
//	router := setupRouter()
//
//	// 构造请求体
//	user := models.User{
//		Username: "testuser",
//		Password: "testpass",
//	}
//
//	body, _ := json.Marshal(user)
//
//	// 创建 POST 请求
//	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
//	req.Header.Set("Content-Type", "application/json")
//
//	// 创建响应记录器
//	w := httptest.NewRecorder()
//
//	// 处理请求
//	router.ServeHTTP(w, req)
//
//	// 断言返回状态码和内容
//	assert.Equal(t, 200, w.Code)
//	assert.Contains(t, w.Body.String(), "注册成功")
//}
