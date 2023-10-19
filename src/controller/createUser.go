package controller

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/validation"
	"github.com/MatheusVict/User-Register-GO/src/controller/model/request"
	"github.com/MatheusVict/User-Register-GO/src/model"
	service2 "github.com/MatheusVict/User-Register-GO/src/model/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	log.Println(userRequest)

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	services := service2.NewUserDomainService()

	if err := services.CreateUser(domain); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.String(http.StatusOK, "UserCreated", domain)
}
