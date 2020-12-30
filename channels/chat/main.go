package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Neil-uli/go-play/channels/chat/models"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var load = godotenv.Load(".env")
var rdb *redis.Client

// WS
// clients is a list of all the currently active clients (or open WebSockets)
var clients = make(map[*websocket.Conn]bool)

// broadcaster is a single channel which is responsible for sending and receiving our ChatMessage data structure
var broadcaster = make(chan models.ChatMessage)

// upgrader is a bit of a clunker; it's necessary to "upgrade" Gorilla incoming requests into a WebSocket connection
var upgrader = websocket.Upgrader{
	// CheckOrigin returns true if the request Origin header is acceptable.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	if load != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Redis client connection
	redisURL := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)

	r := mux.NewRouter()
	// Serving static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	// WebSocket Connection
	r.HandleFunc("/websocket", handleConnections)
	go handleMessages()

	// Listen and start the server
	log.Print("Server starting  at localhost:" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
