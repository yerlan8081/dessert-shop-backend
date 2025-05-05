package routes

import (
	"dessert-shop-backend/microservices/product-service/controllers"
	"dessert-shop-backend/microservices/product-service/middleware"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	cart := r.Group("/api/cart")
	cart.Use(middleware.JWTAuthMiddleware())
	{
		cart.POST("", controllers.AddToCart)
		cart.GET("", controllers.GetCart)
		cart.GET(":id", controllers.GetCartItemByID)
		cart.PUT(":id", controllers.UpdateCartItem)
		cart.DELETE("/:id", controllers.DeleteCartItem)
	}
}
