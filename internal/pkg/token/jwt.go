package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/seagalputra/cashlog/internal/pkg/config"
	"time"
)

var secretKey = config.Get("JWT_SECRET")

func Generate(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := unsignedToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func Parse(token string) (string, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := parse.Claims.(jwt.MapClaims); ok && parse.Valid {
		userID := claims["user_id"].(string)
		return userID, nil
	} else {
		return "", err
	}
}
