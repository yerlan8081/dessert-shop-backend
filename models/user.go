package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"` // "user" or "admin"
}
