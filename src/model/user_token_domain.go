package model

import (
	"fmt"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
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

func VerifyTokenMiddleware(ctx *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(ctx.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, errorsHandle.NewBadRequestError("invalid token")
	})

	if err != nil {
		errorRest := errorsHandle.NewUnauthorizedError("invalid token")
		ctx.JSON(errorRest.Code, errorRest)
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errorRest := errorsHandle.NewUnauthorizedError("invalid token")
		ctx.JSON(errorRest.Code, errorRest)
		ctx.Abort()
		return
	}

	userDomain := &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}
	log.Println("User authenticated ", userDomain)
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	return token
}
