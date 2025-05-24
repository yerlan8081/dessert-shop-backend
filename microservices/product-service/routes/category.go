package routes

import (
	"dessert-shop-backend/microservices/product-service/controllers"
	"dessert-shop-backend/microservices/product-service/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CatRoutes(r *gin.Engine) {
	api := r.Group("/api")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api.Use(middleware.JWTAuthMiddleware())
	{
		// 分类接口
		api.POST("/categories", controllers.CreateCategory)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/categories/:id", controllers.GetCategoryWithDesserts)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)
	}
}
