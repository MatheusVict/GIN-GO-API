package main

import (
	"context"
	"github.com/MatheusVict/User-Register-GO/src/configuration/database/mongodb"
	"github.com/MatheusVict/User-Register-GO/src/controller"
	"github.com/MatheusVict/User-Register-GO/src/controller/routes"
	"github.com/MatheusVict/User-Register-GO/src/model/repository"
	"github.com/MatheusVict/User-Register-GO/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error on .env file")
	}

	database, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatalf("Error trying to connect to database: %s\n", err.Error())
	}
	repos := repository.NewUserRepository(database)

	servic := service.NewUserDomainService(repos)
	userController := controller.NewUserControllerInterface(servic)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln("Error on start application: ", err)
	}
}
