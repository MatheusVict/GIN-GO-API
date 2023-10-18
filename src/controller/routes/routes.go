package routes

import (
	"github.com/MatheusVict/User-Register-GO/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/getUserById/:userId", controller.FindUserById)
	routerGroup.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	routerGroup.POST("/createUser", controller.CreateUser)
	routerGroup.PUT("/updateUser/:userId", controller.UpdateUser)
	routerGroup.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
