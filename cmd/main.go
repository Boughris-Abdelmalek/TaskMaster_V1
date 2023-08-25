package main

import (
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/api"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/postgres"
	"log"
)

func main() {
	store, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", store)
	server.Run()
}
