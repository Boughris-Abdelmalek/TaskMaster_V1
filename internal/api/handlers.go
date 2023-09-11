package api

import (
	"encoding/json"
	"fmt"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/api/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

func HandleLogin(s APIServer, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc, err := s.Store.GetUserByUserName(req.UserName)
	if err != nil {
		return err
	}

	token, err := createJWT(acc)
	if err != nil {
		return err
	}

	resp := models.LoginResponse{
		Token:    token,
		UserName: acc.UserName,
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func HandleUser(s APIServer, w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return HandleGetUser(s, w, r)
	}
	if r.Method == "POST" {
		return HandleCreateUser(s, w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func HandleGetUser(s APIServer, w http.ResponseWriter, r *http.Request) error {
	users, err := s.Store.GetUser()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, users)
}

func HandleCreateUser(s APIServer, w http.ResponseWriter, r *http.Request) error {
	req := new(models.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	user, err := models.NewAccount(req.FirstName, req.LastName, req.UserName, req.Email, req.Password)
	if err != nil {
		return err
	}
	if err := s.Store.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func createJWT(user *models.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"ID":        user.ID,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
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
