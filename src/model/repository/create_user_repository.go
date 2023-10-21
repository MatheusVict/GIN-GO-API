package repository

import (
	"context"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
)


func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("Init createUser repository")
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	ctx := context.Background()

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(ctx, value)
	if err != nil {
		return nil, errorsHandle.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(value), nil
}
