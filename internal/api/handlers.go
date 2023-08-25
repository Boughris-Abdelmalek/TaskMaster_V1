package api

import (
	"encoding/json"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/api/models"
	"log"
	"net/http"
)

// handleGetTodos handles the GET request for fetching todos.
func (s *APIServer) handleGetTodo(w http.ResponseWriter, r *http.Request) error {
	todos, err := s.store.GetTodos()
	if err != nil {
		log.Println("Error getting todos:", err)
		return err
	}

	return WriteJSON(w, http.StatusOK, todos)
}

// handleCreateTodo handles the POST request for creating new todos
func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {
	req := new(models.CreateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	todo, err := models.NewTodo(req.Title, req.Description)
	if err != nil {
		return err
	}
	if err := s.store.CreateTodo(todo); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

// handleGetTodoByID handles the GET request for getting all the todos
func (s *APIServer) handleGetTodoByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	todo, err := s.store.GetTodoByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

// handleDeleteTodo handles the DELETE request for deleting todos
func (s *APIServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteTodo(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// handleUpdateTodo handles the PATCH request for updating the todos
func (s *APIServer) handleUpdateTodo(w http.ResponseWriter, r *http.Request) error {
	req := new(models.UpdateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	id, err := getID(r)
	if err != nil {
		return err
	}

	todo, err := models.NewTodoUpdate(req.Title, req.Description, req.Completed)
	if err != nil {
		return err
	}

	updatedTodo, err := s.store.UpdateTodo(id, todo)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, updatedTodo)
}
