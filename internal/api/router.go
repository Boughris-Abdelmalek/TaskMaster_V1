package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	todosPath  = "/todos"
	todoIDPath = "/todos/{id}"
)

// RegisterRoutes starts the API server and listens for incoming requests.
func (s *APIServer) RegisterRoutes() {
	router := mux.NewRouter()

	router.HandleFunc(todosPath, makeHTTPHandleFunc(HandleGetTodo, s)).Methods("GET")
	router.HandleFunc(todosPath, makeHTTPHandleFunc(HandleCreateTodo, s)).Methods("POST")
	router.HandleFunc(todoIDPath, makeHTTPHandleFunc(HandleGetTodoByID, s)).Methods("GET")
	router.HandleFunc(todoIDPath, makeHTTPHandleFunc(HandleDeleteTodo, s)).Methods("DELETE")
	router.HandleFunc(todoIDPath, makeHTTPHandleFunc(HandleUpdateTodo, s)).Methods("PATCH")

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// apiFunc defines the type of function that handles API requests.
type apiFunc func(s APIServer, w http.ResponseWriter, r *http.Request) error

// makeHTTPHandleFunc converts an API function into a http.HandlerFunc.
func makeHTTPHandleFunc(handlerFunc apiFunc, s *APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(*s, w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
