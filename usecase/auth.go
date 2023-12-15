package usecase

import (
	"errors"
	"kel15/models"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (usecase *Usecase) GetToken(c *gin.Context) (*string, error) {
	headers := c.Request.Header["Authorization"]
	splitToken := strings.Split(headers[0], " ")

	if len(splitToken) != 2 {
		return nil, errors.New("Invalid token")
	}

	token := splitToken[1]

	return &token, nil
}

func (usecase *Usecase) GetUserByToken(c *gin.Context, tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	userId := claims["id"].(float64)

	if err != nil {
		return nil, err
	}

	user, err := usecase.repository.GetUserById(c, int64(userId))
	if err != nil {
		return nil, err
	}

	return user, nil
}
