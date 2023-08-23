package main

import (
	"encoding/json"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/api"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/postgres"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var todoMap map[string]types.Todo

func main() {
	todoMap = make(map[string]types.Todo)

	store, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", store)
	server.Run()

	//router.HandleFunc("/todos", CreateTodo).Methods("POST")
	//router.HandleFunc("/todos/{id}", GetTodoById).Methods("GET")
	//router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")
	//router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
}

func GetTodos(w http.ResponseWriter, r *http.Request, store postgres.PostgresStore) {
	todos, err := store.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo types.Todo
	_ = json.NewDecoder(r.Body).Decode(&newTodo)

	newTodo.ID = uuid.New().String()
	todoMap[newTodo.ID] = newTodo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
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

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todos := []types.Todo{}
	for _, todo := range todoMap {
		if todo.ID != id {
			todos = append(todos, todo)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTodo types.Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)

	updatedTodo.ID = id
	todoMap[id] = updatedTodo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}
