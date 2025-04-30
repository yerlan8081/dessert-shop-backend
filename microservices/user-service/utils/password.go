// utils/password.go
package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	// Hash the password with a cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}
