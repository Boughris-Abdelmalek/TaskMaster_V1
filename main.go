package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var todoMap map[string]Todo

func main() {
	todoMap = make(map[string]Todo)

	router := mux.NewRouter()
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", GetTodosById).Methods("GET")
	router.HandleFunc("/todos/{id}", DeleteTodosById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{}
	for _, todo := range todoMap {
		todos = append(todos, todo)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	_ = json.NewDecoder(r.Body).Decode(&newTodo)

	newTodo.ID = uuid.New().String()
	todoMap[newTodo.ID] = newTodo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

func GetTodosById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, todo := range todoMap {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todo)

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Todo not found"))
}

func DeleteTodosById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todos := []Todo{}
	for _, todo := range todoMap {
		if todo.ID != id {
			todos = append(todos, todo)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
