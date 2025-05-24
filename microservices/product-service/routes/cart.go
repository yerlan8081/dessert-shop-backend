package routes

import (
	"dessert-shop-backend/microservices/product-service/controllers"
	"dessert-shop-backend/microservices/product-service/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CartRoutes(r *gin.Engine) {
	cart := r.Group("/api/cart")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	cart.Use(middleware.JWTAuthMiddleware())
	{
		cart.POST("", controllers.AddToCart)
		cart.GET("", controllers.GetCart)
		cart.GET(":id", controllers.GetCartItemByID)
		cart.PUT(":id", controllers.UpdateCartItem)
		cart.DELETE("/:id", controllers.DeleteCartItem)
	}
}
