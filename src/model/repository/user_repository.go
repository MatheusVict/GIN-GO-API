package repository

import (
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *errorsHandle.ErrorsHandle)
	FindUserByEmail(email string) (model.UserDomainInterface, *errorsHandle.ErrorsHandle)
	FindUserByID(id string) (model.UserDomainInterface, *errorsHandle.ErrorsHandle)
}
