package authDevice

import (
	"github.com/gorilla/websocket"
	"log"
)

func Auth(json *map[string]interface{}, messageType int, conn *websocket.Conn) error {
	if value, ok := (*json)["OS"]; ok {
		log.Println("OS: " + value.(string))
	}

	if volumesInfo, ok := (*json)["Volumes Info"]; ok {
		if volumes, ok := volumesInfo.(map[string]interface{})["Volume serial numbers"]; ok {
			for _, volume := range volumes.([]interface{}) {
				log.Println("volume serial number: " + volume.(string))
			}
		}
	}

	if biosInfo, ok := (*json)["BIOS Info"]; ok {
		for key, value := range biosInfo.(map[string]interface{}) {
			if key == "Some data" {
				for _, data := range value.([]interface{}) {
					log.Println("Some data: " + data.(string))
				}
			} else {
				log.Println(key + ": " + value.(string))
			}
		}
	}

	message := []byte("Ok")

	if err := conn.WriteMessage(messageType, message); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
