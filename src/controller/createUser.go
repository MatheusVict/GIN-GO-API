package controller

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/validation"
	"github.com/MatheusVict/User-Register-GO/src/controller/model/request"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (user *userControllerInterface) CreateUser(ctx *gin.Context) {
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

	if err := user.service.CreateUser(domain); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
