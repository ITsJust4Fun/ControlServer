package database

import (
	"ControlServer/pkg/config"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
)

func setFieldToInterface(input *interface{}, fieldName string, value interface{}) {
	valueInterface := reflect.ValueOf(&input).Elem()
	tmp := reflect.New(valueInterface.Elem().Type()).Elem()
	tmp.Set(valueInterface.Elem())
	tmp.FieldByName(fieldName).Set(reflect.ValueOf(value))
	valueInterface.Set(tmp)
}

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

	setFieldToInterface(&input, "ID", primitive.NewObjectID())

	_, err = collection.InsertOne(ctx, input)

	if err != nil {
		log.Print("Error when inserting", err)
		return "", err
	}

	return input, nil
}
