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

func (user *userControllerInterface) LoginUser(ctx *gin.Context) {
	log.Println("LoginUserController init")
	var userRequest request.UserLogin

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		log.Println("Error on validate user")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	log.Println(userRequest)

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, err := user.service.LoginUserService(domain)
	if err != nil {
		log.Println("error on login user in service to controller")
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
