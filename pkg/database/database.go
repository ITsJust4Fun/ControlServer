package database

import (
	"ControlServer/pkg/config"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
)

func InsertOne(input interface{}, collectionName string) (interface{}, error) {
	var empty interface{}
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
		return empty, err
	}

	collection := client.Database("control_server").Collection(collectionName)
	err = client.Connect(ctx)

	if err != nil {
		log.Panic("Error when connecting to mongodb", err)
		return empty, err
	}

	id := reflect.ValueOf(&input).Elem().Elem().FieldByName("ID")

	if id.IsValid() {
		id.Set(reflect.ValueOf(primitive.NewObjectID()))
	} else {
		return empty, errors.New("No id field!")
	}

	_, err = collection.InsertOne(ctx, input)

	if err != nil {
		log.Print("Error when inserting", err)
		return "", err
	}

	return input, nil
}
