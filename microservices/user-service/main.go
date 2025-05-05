package main

import (
	"dessert-shop-backend/microservices/user-service/database"
	"dessert-shop-backend/microservices/user-service/models"
	"dessert-shop-backend/microservices/user-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()

	// 自动迁移所有模型
	database.DB.AutoMigrate(
		&models.User{},
	)

	routes.AuthRoutes(r)
	r.Run(":8081")
}
