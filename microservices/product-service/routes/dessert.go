package routes

import (
	"dessert-shop-backend/microservices/product-service/controllers"
	"dessert-shop-backend/microservices/product-service/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func DessertRoutes(r *gin.Engine) {
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
		api.POST("/desserts", controllers.CreateDessert)
		api.GET("/desserts", controllers.GetDesserts)
		api.GET("/desserts/:id", controllers.GetDessert)
		api.PUT("/desserts/:id", controllers.UpdateDessert)
		api.DELETE("/desserts/:id", controllers.DeleteDessert)
		api.GET("desserts/category/:id", controllers.GetDessertsByCategory)
	}
}
