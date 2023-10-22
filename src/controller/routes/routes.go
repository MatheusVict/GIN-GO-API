package routes

import (
	"github.com/MatheusVict/User-Register-GO/src/controller"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	routerGroup *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {

	routerGroup.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	routerGroup.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	routerGroup.POST("/createUser", userController.CreateUser)
	routerGroup.PUT("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	routerGroup.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)
	routerGroup.POST("/login", userController.LoginUser)
}
