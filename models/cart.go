package models

import "gorm.io/gorm"

// 购物车项
type CartItem struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	DessertID uint    `json:"dessert_id"`
	Quantity  int     `json:"quantity"`
	Dessert   Dessert ` json:"dessert" gorm:"foreignKey:DessertID"` //json:"-"
}
