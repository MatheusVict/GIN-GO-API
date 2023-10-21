package controller

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/configuration/validation"
	"github.com/MatheusVict/User-Register-GO/src/controller/model/request"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
)

func (user *userControllerInterface) UpdateUser(ctx *gin.Context) {
	log.Println("UpdateUserController init")
	var userRequest request.UserUpdateRequest

	userId := ctx.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil ||
		strings.TrimSpace(userId) == "" {
		errorMessage := errorsHandle.NewBadRequestError("UserID is not a valid ID")
		log.Println("Error on parse to uuid ", err)

		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		log.Println("Error on validate user")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	log.Println(userRequest)

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := user.service.UpdateUser(userId, domain)
	if err != nil {
		log.Println("error on update user in service to controller :", userId)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.Status(http.StatusOK)
}
