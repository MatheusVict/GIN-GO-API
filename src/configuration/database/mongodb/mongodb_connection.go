package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const (
	MONGODB_URI      = "MONGODB_CONNECTION_URI"
	MONGODB_DATABASE = "MONGODB_USER_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	uriConnection := os.Getenv(MONGODB_URI)
	mongodbDatabase := os.Getenv(MONGODB_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uriConnection))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected")

	return client.Database(mongodbDatabase), nil
}
