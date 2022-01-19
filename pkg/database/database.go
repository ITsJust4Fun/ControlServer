package database

import (
	"ControlServer/pkg/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InsertOne(input interface{}, collectionName string) (string, error) {
	conf := config.GetConfig()
	ctx, cancel := context.WithTimeout(context.Background(), conf.DatabaseTimeout)
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.DatabaseURI))

	defer func(client *mongo.Client, ctx context.Context) {
		err = client.Disconnect(ctx)

		if err != nil {
			log.Panic("Error when disconnecting to mongodb", err)
		}

		cancel()
	}(client, ctx)

	if err != nil {
		log.Panic("Error when creating mongodb connection client", err)
		return "", err
	}

	collection := client.Database("control_server").Collection(collectionName)
	err = client.Connect(ctx)

	if err != nil {
		log.Panic("Error when connecting to mongodb", err)
	}

	id, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Print("Error when inserting", err)
		return "", err
	}

	return id.InsertedID.(string), nil
}
