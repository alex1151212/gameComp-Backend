package utils

import (
	"errors"
	"fmt"
	models "gameComp-Backend/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("")

type AuthClaims struct {
	UserId   uint   `json:"userId"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsUpload bool   `json:"isUpload"`
	jwt.RegisteredClaims
}

func InitJWTKey() {
	jwtKey = []byte(os.Getenv("JWT_KEY"))
}

func GenerateToken(user models.User) (string, error) {
	claims := AuthClaims{
		user.ID,
		user.Email,
		user.Phone,
		user.Team.IsUpload,
		jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
			Issuer:    "IGD",
		},
	}
	// 創建一個新的令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 簽署令牌
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*models.User, error) {
	var claims AuthClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return &models.User{}, err
	}
	if !token.Valid {
		return &models.User{}, errors.New("invalid token")
	}

	user := models.User{
		Email: claims.Email,
	}

	user.FindOne()
	return &user, nil
}
