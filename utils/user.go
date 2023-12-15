package utils

import (
	"fmt"
	"kel15/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(payloadPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(payloadPassword))
	if err != nil {
		return err
	}
	return nil
}

func GenerateToken(user *models.User) (string, error) {
	key := os.Getenv("JWT_KEY")

	fmt.Println(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    &user.ID,
		"email": &user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
