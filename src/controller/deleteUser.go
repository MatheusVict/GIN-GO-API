package controller

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
)

func (user *userControllerInterface) DeleteUser(ctx *gin.Context) {
	log.Println("DeleteUserController init")

	userId := ctx.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil ||
		strings.TrimSpace(userId) == "" {
		errorMessage := errorsHandle.NewBadRequestError("UserID is not a valid ID")
		log.Println("Error on parse to uuid ", err)

		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	err := user.service.DeleteUser(userId)
	if err != nil {
		log.Println("error on update user in service to controller :", userId)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.Status(http.StatusOK)
}
