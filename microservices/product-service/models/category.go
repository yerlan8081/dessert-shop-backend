package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `json:"name" binding:"required"`
	Desserts []Dessert `json:"desserts"  gorm:"foreignKey:CategoryID"`
}
