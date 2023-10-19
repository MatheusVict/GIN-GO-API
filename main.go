package main

import (
	"fmt"
	"github.com/MatheusVict/User-Register-GO/src/controller"
	"github.com/MatheusVict/User-Register-GO/src/controller/routes"
	"github.com/MatheusVict/User-Register-GO/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error on .env file")
	}
	fmt.Println(os.Getenv("TEST"))

	servic := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(servic)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln("Error on start application: ", err)
	}
}
