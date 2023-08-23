package api

import (
	"encoding/json"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/postgres"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/types"
	"github.com/gorilla/mux"
	"log"
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

// Run starts the API server and listens for incoming requests.
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", makeHTTPHandleFunc(s.handleGetTodo)).Methods("GET")
	router.HandleFunc("/todos", makeHTTPHandleFunc(s.handleCreateTodo)).Methods("POST")
	router.HandleFunc("/todos/{id}", makeHTTPHandleFunc(s.handleGetTodoByID)).Methods("GET")

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// handleGetTodos handles the GET request for fetching todos.
func (s *APIServer) handleGetTodo(w http.ResponseWriter, r *http.Request) error {
	todos, err := s.store.GetTodos()
	if err != nil {
		log.Println("Error getting todos:", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, todos)
}

func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	todo, err := types.NewTodo(req.Title, req.Description)
	if err != nil {
		return err
	}
	if err := s.store.CreateTodo(todo); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

func (s *APIServer) handleGetTodoByID(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	todo, err := s.store.GetTodoByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
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
