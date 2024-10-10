package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/VivekHalder/TryingDocs/database"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Document struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

var documents = make(map[string]Document)
var clients = make(map[*websocket.Conn]bool)
var clientMutex = &sync.Mutex{}
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleClient(conn *websocket.Conn) {
	defer conn.Close()

	clientMutex.Lock()
	clients[conn] = true
	clientMutex.Unlock()

	// Send initial documents to the new client
	for _, doc := range documents {
		if err := conn.WriteJSON(doc); err != nil {
			log.Println("Error sending initial document:", err)
			clientMutex.Lock()
			delete(clients, conn)
			clientMutex.Unlock()
			return
		}
	}

	for {
		var msg map[string]interface{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			clientMutex.Lock()
			delete(clients, conn)
			clientMutex.Unlock()
			break
		}

		var doc Document
		docMap, ok := msg["data"].(map[string]interface{})
		if !ok {
			log.Println("Invalid document data:", msg)
			continue
		}
		docBytes, err := json.Marshal(docMap)
		if err != nil {
			log.Println("Error marshalling document data:", err)
			continue
		}
		if err := json.Unmarshal(docBytes, &doc); err != nil {
			log.Println("Error unmarshalling document:", err)
			continue
		}
		documents[doc.ID] = doc
		clientMutex.Lock()
		for client := range clients {
			if err := client.WriteJSON(doc); err != nil {
				log.Println("Error sending document to client:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientMutex.Unlock()

	}
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	go handleClient(conn)
}

func main() {

	database.InitDB()
	defer database.CloseDB()

	fmt.Println("Server started!")
	r := mux.NewRouter()
	r.HandleFunc("/ws", HandleConnection)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
