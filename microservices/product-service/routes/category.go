package routes

import (
	"dessert-shop-backend/microservices/product-service/controllers"
	"dessert-shop-backend/microservices/product-service/middleware"
	"github.com/gin-gonic/gin"
)

func CatRoutes(r *gin.Engine) {
	api := r.Group("/api")
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
