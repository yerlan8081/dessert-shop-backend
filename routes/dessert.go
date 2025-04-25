package routes

import (
	"dessert-shop-backend/controllers"
	"dessert-shop-backend/middleware"
	"github.com/gin-gonic/gin"
)

func DessertRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.POST("/desserts", controllers.CreateDessert)
		api.GET("/desserts", controllers.GetDesserts)
		api.GET("/desserts/:id", controllers.GetDessert)
		api.PUT("/desserts/:id", controllers.UpdateDessert)
		api.DELETE("/desserts/:id", controllers.DeleteDessert)
		api.GET("desserts/category/:id", controllers.GetDessertsByCategory)
	}
}
