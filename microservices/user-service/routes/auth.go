package routes

import (
	"dessert-shop-backend/microservices/product-service/middleware"
	"dessert-shop-backend/microservices/user-service/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func AuthRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.PUT("/user", middleware.JWTAuthMiddleware(), controllers.UpdateUser)
}
