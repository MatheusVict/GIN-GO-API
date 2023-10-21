package model

import (
	"fmt"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/golang-jwt/jwt"
	"os"
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
