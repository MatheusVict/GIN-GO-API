package routes

import (
	"github.com/MatheusVict/User-Register-GO/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	routerGroup *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {

	routerGroup.GET("/getUserById/:userId", userController.FindUserByID)
	routerGroup.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	routerGroup.POST("/createUser", userController.CreateUser)
	routerGroup.PUT("/updateUser/:userId", userController.UpdateUser)
	routerGroup.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
