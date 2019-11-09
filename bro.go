package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello World"); err != nil {
		log.Fatalln(err)
	}
}

type Message struct {
	Message string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		if err := ws.ReadJSON(&msg); err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	router := mux.NewRouter()
	buildHandler := http.FileServer(http.Dir("frontend/build"))
	router.PathPrefix("/").Handler(buildHandler)

	http.HandleFunc("/ws", handleConnections)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln(err)
	}
	go handleMessages()
}

