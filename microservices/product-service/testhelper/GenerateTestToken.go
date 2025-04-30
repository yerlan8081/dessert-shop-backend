package testhelper

import (
	"dessert-shop-backend/microservices/user-service/models"
	"dessert-shop-backend/microservices/user-service/utils"
)

func GenerateTestToken(user models.User) string {
	token, _ := utils.GenerateJWT(user.ID, user.Username, user.Role)
	return token
}
