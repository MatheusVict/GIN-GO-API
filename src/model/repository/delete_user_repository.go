package repository

import (
	"context"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
)

func (user *userRepository) DeleteUser(userID string) *errorsHandle.ErrorsHandle {
	log.Println("Init DeleteUser repository")
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := user.databaseConnection.Collection(collectionName)
	userIDHex, _ := primitive.ObjectIDFromHex(userID)

	ctx := context.Background()
	filter := bson.D{{Key: "_id", Value: userIDHex}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error on delete user: ", err)
		return errorsHandle.NewInternalServerError(err.Error())
	}

	log.Println("User ID: ", userID)

	return nil
}
