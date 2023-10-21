package repository

import (
	"context"
	"fmt"
	"github.com/MatheusVict/User-Register-GO/src/configuration/errorsHandle"
	"github.com/MatheusVict/User-Register-GO/src/model"
	"github.com/MatheusVict/User-Register-GO/src/model/repository/entity"
	"github.com/MatheusVict/User-Register-GO/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("Init findUserByEmail repository")
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	ctx := context.Background()
	filter := bson.D{{"email", email}}

	err := collection.FindOne(ctx, filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			log.Println("User not found: ", errorMessage)
			return nil, errorsHandle.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		log.Println("Error on search an user: ", err)
		return nil, errorsHandle.NewInternalServerError(errorMessage)
	}

	log.Println("User find with successfully")
	log.Println("User email: ", userEntity.Email)
	return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("Init findUserByID repository")
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	ctx := context.Background()
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}

	err := collection.FindOne(ctx, filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this ID: %s", id)
			log.Println("User not found: ", errorMessage)
			return nil, errorsHandle.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by ID"
		log.Println("Error on search an user: ", err)
		return nil, errorsHandle.NewInternalServerError(errorMessage)
	}

	log.Println("User find with successfully")
	log.Println("User email: ", userEntity.ID.Hex())
	return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (model.UserDomainInterface, *errorsHandle.ErrorsHandle) {
	log.Println("Init findUserByEmailAndPassword repository")
	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	ctx := context.Background()
	filter := bson.D{
		{"email", email},
		{"password", password},
	}

	err := collection.FindOne(ctx, filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Email or password invalid"
			log.Println("User not found: ", errorMessage)
			return nil, errorsHandle.NewForbiddenError(errorMessage)
		}

		errorMessage := "Error trying to find user by email and password"
		log.Println("Error on search an user: ", err)
		return nil, errorsHandle.NewInternalServerError(errorMessage)
	}

	log.Println("User find with successfully")
	log.Println("User email: ", userEntity.Email)
	return converter.ConvertEntityToDomain(userEntity), nil
}
