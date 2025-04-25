package routes

import (
	"dessert-shop-backend/controllers"
	"dessert-shop-backend/middleware"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	cart := r.Group("/api/cart")
	cart.Use(middleware.JWTAuthMiddleware())
	{
		cart.POST("", controllers.AddToCart)
		cart.GET("", controllers.GetCart)
		cart.DELETE("/:id", controllers.DeleteCartItem)
	}
}
