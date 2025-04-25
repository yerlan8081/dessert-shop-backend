package main

import (
	"dessert-shop-backend/database"
	"dessert-shop-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()

	// 自动迁移所有模型
	//database.DB.AutoMigrate(
	//	&models.Category{},
	//	&models.User{},
	//	&models.Dessert{},
	//	&models.Order{},
	//	&models.OrderItem{},
	//	&models.CartItem{},
	//)

	routes.AuthRoutes(r)
	routes.DessertRoutes(r)
	routes.CatRoutes(r)
	routes.CartRoutes(r)

	r.Run(":8080")
}
