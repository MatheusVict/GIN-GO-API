package controller

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"net/mail"
)

func (user *userControllerInterface) FindUserByID(ctx *gin.Context) {
	log.Println("Init FindUserByID on controller")

	userId := ctx.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := errorsHandle.NewBadRequestError("UserID is not a valid ID")
		log.Println("Error on parse to uuid ", err)

		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := user.service.FindUserByIDService(userId)
	if err != nil {
		log.Println("Error to find user")
		ctx.JSON(err.Code, err)
	}

	convertedToResponse := view.ConvertDomainToResponse(userDomain)
	ctx.JSON(http.StatusOK, convertedToResponse)
}

func (user *userControllerInterface) FindUserByEmail(ctx *gin.Context) {
	log.Println("Init FindUserByEmail on controller")

	userEmail := ctx.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := errorsHandle.NewBadRequestError("Email is not a valid email")
		log.Println("Error on parse Email ", err)

		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := user.service.FindUserByEmailService(userEmail)
	if err != nil {
		log.Println("Error to find user")
		ctx.JSON(err.Code, err)
	}

	convertedToResponse := view.ConvertDomainToResponse(userDomain)
	ctx.JSON(http.StatusOK, convertedToResponse)
}
