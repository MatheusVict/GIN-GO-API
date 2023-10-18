package controller

import (
	"fmt"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := errorsHandle.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()),
		)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	log.Println(userRequest)
}
