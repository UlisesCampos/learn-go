package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var load = godotenv.Load(".env")

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	if load != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
