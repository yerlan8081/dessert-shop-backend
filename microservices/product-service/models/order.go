package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	DessertID uint    `json:"dessert_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // 单价
}
