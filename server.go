package main

import (
	"ControlServer/graph"
	"ControlServer/graph/generated"
	"ControlServer/internal/device"
	"ControlServer/pkg/config"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func readFromWebSocket(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, messageBytes, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		var result map[string]interface{}
		err = json.Unmarshal(messageBytes, &result)

		if err != nil {
			log.Println(err)
			return
		}

		method, ok := result["method"]

		if !ok {
			log.Println("no method field")
			return
		}

		switch method {
		case "auth":
			err := device.Auth(messageBytes, messageType, conn)

			if err != nil {
				log.Println(err)
				return
			}

			break
		default:
			message := []byte("unknown method")

			if err := conn.WriteMessage(messageType, message); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func sendError(conn *websocket.Conn) {

}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	readFromWebSocket(ws)
}

func main() {
	config.ReadConfigFile()
	conf := config.GetConfig()

	port := os.Getenv("PORT")
	if port == "" {
		port = conf.Port
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/ws", wsEndpoint)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
