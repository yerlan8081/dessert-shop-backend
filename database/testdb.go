package database

import (
	"fmt"
	"log"

	"dessert-shop-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func InitTestDB() *gorm.DB {
	// 直接在代码中指定连接信息
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "20021512" // ← 请替换为你自己的 PostgreSQL 密码
	dbname := "db_test"    // ← 测试数据库名

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ 无法连接测试数据库: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Dessert{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{}) // 加你需要测试的模型
	if err != nil {
		log.Fatalf("❌ 自动迁移失败: %v", err)
	}

	TestDB = db
	return db
}
