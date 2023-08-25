package api

import (
	"encoding/json"
	"fmt"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/postgres"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// APIServer represents the API server structure
type APIServer struct {
	listenAddr string
	store      postgres.Storage
}

// NewAPIServer creates a new instance of APIServer
func NewAPIServer(listenAddr string, store postgres.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

// WriteJSON is a Helper function to return JSON response to client
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// apiFunc defines the type of function that handles API requests.
type apiFunc func(w http.ResponseWriter, r *http.Request) error

// ApiError represents an error response in JSON format.
type ApiError struct {
	Error string `json:"error"`
}

// makeHTTPHandleFunc converts an API function into a http.HandleFunc.
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// getID is a helper function that returns the id from the url
func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
