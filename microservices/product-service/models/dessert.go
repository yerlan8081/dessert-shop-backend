package models

import "gorm.io/gorm"

type Dessert struct {
	gorm.Model
	Name        string  `json:"name" binding:"required,min=2,max=100"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	CategoryID  uint    `json:"category_id"`
}
