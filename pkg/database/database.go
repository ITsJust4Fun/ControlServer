package database

import (
	"ControlServer/pkg/config"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
)

type ConnectionControl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func SetFieldToInterface(input interface{}, fieldName string, value interface{}) interface{} {
	if reflect.ValueOf(input).Kind() == reflect.Struct {
		valueInterface := reflect.ValueOf(&input).Elem()
		tmp := reflect.New(valueInterface.Elem().Type()).Elem()
		tmp.Set(valueInterface.Elem())
		tmp.FieldByName(fieldName).Set(reflect.ValueOf(value))

		return tmp.Interface()
	}

	valueInterface := reflect.ValueOf(&input).Elem()
	tmp := reflect.New(valueInterface.Elem().Elem().Type()).Elem()
	tmp.Set(valueInterface.Elem().Elem())
	tmp.FieldByName(fieldName).Set(reflect.ValueOf(value))
	CloneValueToPointer(tmp.Interface(), input)

	return tmp.Interface()
}

func CloneValueToPointer(source interface{}, destination interface{}) {
	x := reflect.ValueOf(source)
	starX := x
	y := reflect.New(starX.Type())
	starY := y.Elem()
	starY.Set(starX)
	reflect.ValueOf(destination).Elem().Set(y.Elem())
}

func ClonePointerToPointer(source interface{}, destination interface{}) {
	x := reflect.ValueOf(source)
	starX := x.Elem()
	y := reflect.New(starX.Type())
	starY := y.Elem()
	starY.Set(starX)
	reflect.ValueOf(destination).Elem().Set(y.Elem())
}

func connect(collectionName string) (*ConnectionControl, func(), error) {
	conf := config.GetConfig()
	ctx, cancel := context.WithTimeout(context.Background(), conf.DatabaseTimeout)
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.DatabaseURI))

	disconnect := func() {
		_ = client.Disconnect(ctx)
		cancel()
	}

	if err != nil {
		log.Println("Error when creating mongodb connection client", err)
		disconnect()
		return nil, func() {}, err
	}

	collection := client.Database(conf.DatabaseName).Collection(collectionName)
	err = client.Connect(ctx)

	if err != nil {
		log.Println("Error when connecting to mongodb", err)
		disconnect()
		return nil, func() {}, err
	}

	return &ConnectionControl{collection: collection, ctx: ctx}, disconnect, nil
}

func Ping() bool {
	conf := config.GetConfig()
	ctx, cancel := context.WithTimeout(context.Background(), conf.DatabaseTimeout)
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.DatabaseURI))

	disconnect := func() {
		_ = client.Disconnect(ctx)
		cancel()
	}

	if err != nil {
		log.Println("Error when creating mongodb connection client", err)
		disconnect()
		return false
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Println("Error when connecting to mongodb", err)
		disconnect()
		return false
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Println(err)
		disconnect()
		return false
	}

	return true
}

func CreateNewDocument(input interface{}, collectionName string) error {
	collectionControl, disconnect, err := connect(collectionName)
	defer disconnect()

	if err != nil {
		return err
	}

	inputWithId := SetFieldToInterface(input, "ID", primitive.NewObjectID())
	_, err = collectionControl.collection.InsertOne(collectionControl.ctx, inputWithId)

	if err != nil {
		log.Print("Error when inserting", err)
		return err
	}

	return nil
}

func UpdateOne(input interface{}, update interface{}, collectionName string) error {
	collectionControl, disconnect, err := connect(collectionName)
	defer disconnect()

	if err != nil {
		return err
	}

	idValue := reflect.ValueOf(input).Elem().FieldByName("ID")
	id := idValue.Interface().(primitive.ObjectID)

	_, err = collectionControl.collection.UpdateOne(collectionControl.ctx, bson.M{"_id": id}, update)

	if err != nil {
		log.Print("Error when updating", err)
		return err
	}

	return nil
}

func FindOne(input interface{}, filter interface{}, collectionName string) error {
	collectionControl, disconnect, err := connect(collectionName)
	defer disconnect()

	if err != nil {
		return err
	}

	err = collectionControl.collection.FindOne(collectionControl.ctx, filter).Decode(input)

	if err != nil {
		return err
	}

	return nil
}

func FindMany(input interface{}, filter interface{}, collectionName string) error {
	collectionControl, disconnect, err := connect(collectionName)
	defer disconnect()

	if err != nil {
		return err
	}

	result, err := collectionControl.collection.Find(collectionControl.ctx, filter)

	if err != nil {
		log.Print("Error when find", err)
		return err
	}

	defer func(result *mongo.Cursor, ctx context.Context) {
		_ = result.Close(ctx)
	}(result, collectionControl.ctx)

	err = result.All(collectionControl.ctx, input)

	if err != nil {
		log.Print("Error when reading reports from cursor", err)
		return err
	}

	return nil
}

func GetAll(input interface{}, collectionName string) error {
	collectionControl, disconnect, err := connect(collectionName)
	defer disconnect()

	if err != nil {
		return err
	}

	result, err := collectionControl.collection.Find(collectionControl.ctx, bson.D{})

	if err != nil {
		log.Print("Error when find", err)
		return err
	}

	defer func(result *mongo.Cursor, ctx context.Context) {
		_ = result.Close(ctx)
	}(result, collectionControl.ctx)

	err = result.All(collectionControl.ctx, input)

	if err != nil {
		log.Print("Error when reading reports from cursor", err)
		return err
	}

	return nil
}

func RemoveOne(input interface{}, collectionName string) error {
	collectionControl, disconnect, err := connect(collectionName)
	defer disconnect()

	if err != nil {
		return err
	}

	idValue := reflect.ValueOf(input).Elem().FieldByName("ID")
	id := idValue.Interface().(primitive.ObjectID)

	result, err := collectionControl.collection.DeleteOne(
		collectionControl.ctx,
		bson.M{"_id": id},
	)

	if err != nil {
		log.Print("Error when removing", err)
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("Can't find item to remove")
	}

	return nil
}

func RemoveByFilter(filter interface{}, collectionName string) (int64, error) {
	collectionControl, disconnect, err := connect(collectionName)
	defer disconnect()

	if err != nil {
		return 0, err
	}

	result, err := collectionControl.collection.DeleteMany(collectionControl.ctx, filter)

	if err != nil {
		log.Print("Error when removing", err)
		return 0, err
	}

	return result.DeletedCount, nil
}
