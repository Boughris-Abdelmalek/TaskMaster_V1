package api

import (
	"encoding/json"
	"fmt"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/api/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// HandleGetTodo handles the GET request for fetching todos.
func HandleGetTodo(s APIServer, w http.ResponseWriter, r *http.Request) error {
	todos, err := s.Store.GetTodos()
	if err != nil {
		log.Println("Error getting todos:", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, todos)
}

// HandleCreateTodo handles the POST request for creating new todos
func HandleCreateTodo(s APIServer, w http.ResponseWriter, r *http.Request) error {
	req := new(models.CreateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	todo, err := models.NewTodo(req.Title, req.Description)
	if err != nil {
		return err
	}
	if err := s.Store.CreateTodo(todo); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

// HandleGetTodoByID handles the GET request for getting all the todos
func HandleGetTodoByID(s APIServer, w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(r)
	if err != nil {
		return err
	}

	todo, err := s.Store.GetTodoByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

// HandleDeleteTodo handles the DELETE request for deleting todos
func HandleDeleteTodo(s APIServer, w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(r)
	if err != nil {
		return err
	}

	if err := s.Store.DeleteTodo(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// HandleUpdateTodo handles the PATCH request for updating the todos
func HandleUpdateTodo(s APIServer, w http.ResponseWriter, r *http.Request) error {
	req := new(models.UpdateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	id, err := GetID(r)
	if err != nil {
		return err
	}

	todo, err := models.NewTodoUpdate(req.Title, req.Description, req.Completed)
	if err != nil {
		return err
	}

	updatedTodo, err := s.Store.UpdateTodo(id, todo)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, updatedTodo)
}

// WriteJSON is a Helper function to return JSON response to ui
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// ApiError represents an error response in JSON format.
type ApiError struct {
	Error string `json:"error"`
}

// GetID is a helper function that returns the id from the url
func GetID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
