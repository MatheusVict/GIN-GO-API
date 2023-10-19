package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitConnection(uriConnection string) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uriConnection))

	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	log.Println("Connected")
}
