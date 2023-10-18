package controller

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	err := errorsHandle.NewBadRequestError("Wrong path")
	ctx.JSON(err.Code, err)
}
