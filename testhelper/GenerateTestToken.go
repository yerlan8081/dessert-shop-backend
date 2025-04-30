package testhelper

import (
	"dessert-shop-backend/models"
	"dessert-shop-backend/utils"
)

func GenerateTestToken(user models.User) string {
	token, _ := utils.GenerateJWT(user.ID, user.Username, user.Role)
	return token
}
