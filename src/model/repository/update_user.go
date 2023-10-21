package repository

import (
	"context"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
)

func (ur *userRepository) UpdateUser(id string, userDomain model.UserDomainInterface) *errorsHandle.ErrorsHandle {
	log.Println("Init UpdateUser repository")
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)
	userIDHex, _ := primitive.ObjectIDFromHex(id)

	ctx := context.Background()
	value := converter.ConvertDomainToEntity(userDomain)
	filter := bson.D{{Key: "_id", Value: userIDHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error on update user: ", err)
		return errorsHandle.NewInternalServerError(err.Error())
	}

	log.Println("User ID: ", id)

	return nil
}
