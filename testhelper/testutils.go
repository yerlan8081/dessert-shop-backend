package testhelper

import (
	"dessert-shop-backend/database"
	"gorm.io/gorm"
	"os"
	"testing"
)

func CleanDatabase(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	db.Exec("TRUNCATE TABLE desserts RESTART IDENTITY CASCADE")
	db.Exec("TRUNCATE TABLE cart_items RESTART IDENTITY CASCADE")
	db.Exec("TRUNCATE TABLE categories RESTART IDENTITY CASCADE")
	db.Exec("TRUNCATE TABLE order_items RESTART IDENTITY CASCADE")
	//db.Exec("TRUNCATE TABLE order RESTART IDENTITY CASCADE")
}

func SetupTestMain(m *testing.M) {
	// 可以添加公共环境变量配置
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "your_password")
	os.Setenv("DB_NAME", "dessert_shop_test")

	database.InitTestDB()
	os.Exit(m.Run())
}
