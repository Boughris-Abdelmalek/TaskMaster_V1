package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Run starts the API server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", makeHTTPHandleFunc(s.handleGetTodo)).Methods("GET")
	router.HandleFunc("/todos", makeHTTPHandleFunc(s.handleCreateTodo)).Methods("POST")
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(s.handleGetTodoByID)).Methods("GET")
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(s.handleDeleteTodo)).Methods("DELETE")
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(s.handleUpdateTodo)).Methods("PATCH")

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
