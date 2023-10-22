package model

import (
	"fmt"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
	"time"
)

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *errorsHandle.ErrorsHandle) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errorsHandle.NewInternalServerError(fmt.Sprintf("Error trying to generate jwt token, err=%s", err))
	}
	return tokenString, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *errorsHandle.ErrorsHandle) {
	secret := os.Getenv(JWT_SECRET_KEY)
	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, errorsHandle.NewBadRequestError("invalid token")
	})

	if err != nil {
		return nil, errorsHandle.NewUnauthorizedError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errorsHandle.NewUnauthorizedError("invalid token")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	return token
}
