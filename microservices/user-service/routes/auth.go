package routes

import (
	"dessert-shop-backend/microservices/product-service/middleware"
	"dessert-shop-backend/microservices/user-service/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.PUT("/user", middleware.JWTAuthMiddleware(), controllers.UpdateUser)
}
