package main

import (
	"github.com/MatheusVict/User-Register-GO/src/controller"
	"github.com/MatheusVict/User-Register-GO/src/model/repository"
	"github.com/MatheusVict/User-Register-GO/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repos := repository.NewUserRepository(database)

	servic := service.NewUserDomainService(repos)
	return controller.NewUserControllerInterface(servic)
}
