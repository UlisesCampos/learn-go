package main

import (
	"net/http"

	"github.com/gorilla/pat"
)

// ServeMux node
var mux *pat.Router = pat.New()

func init() {
	mux.Get("/notes", noteHandler)
	mux.Get("/", helloHandler)
}

func main() {
	err := http.ListenAndServe(address, mux)
}
