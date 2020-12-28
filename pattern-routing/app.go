package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/rs/cors"
)

// ServeMux node
var mux *pat.Router = pat.New()

// init registers patterns and handlers on mux
func init() {
	// api
	mux.Post("/api/v1/notes", noteCreate)
	mux.Get("/api/v1/notes/{id}", noteList)
	mux.Put("/api/v1/notes/{id}", noteUpdate)
	mux.Delete("/api/v1/notes/{id}", noteDelete)
	// login
	mux.Get("/signup", signupHandler)
	mux.Get("/login", loginHandler)
	mux.Get("/logout", logoutHandler)
	// frontend
	mux.Get("/notes/{id}", noteDetailHandler)
	mux.Get("/profile", profileHandler)
	// default
	mux.Get("/", defaultHandler)
}

func main() {
	handler := cors.AllowAll().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "signup")
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login")
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout")
}
func profileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "profile")
}
func noteDetailHandler(w http.ResponseWriter, r *http.Request) {
	noteId := r.URL.Query().Get(":id")
	fmt.Fprintf(w, "note %s detail", noteId)
}
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// API
func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "note list")
}
func noteCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "note create")
}
func noteUpdate(w http.ResponseWriter, r *http.Request) {
	noteId := r.URL.Query().Get(":id")
	fmt.Fprintf(w, " update note %s", noteId)
}
func noteDelete(w http.ResponseWriter, r *http.Request) {
	noteId := r.URL.Query().Get(":id")
	fmt.Fprintf(w, "delete note %s", noteId)
}
