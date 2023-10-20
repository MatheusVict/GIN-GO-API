package repository

import (
	"context"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"log"
	"os"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("Init createUser repository")
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	ctx := context.Background()

	value, err := userDomain.GetJSONValue()

	if err != nil {
		return nil, errorsHandle.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(ctx, value)
	if err != nil {
		return nil, errorsHandle.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
