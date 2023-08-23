package api

import (
	"encoding/json"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/postgres"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
	store      postgres.Storage
}

func NewAPIServer(listenAddr string, store postgres.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", makeHTTPHandleFunc(s.handleGetTodos))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetTodos(w http.ResponseWriter, r *http.Request) error {
	todos, err := s.store.GetTodos()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todos)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
