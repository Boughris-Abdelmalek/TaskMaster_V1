package main

import (
	"encoding/json"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/api"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/postgres"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/types"
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
