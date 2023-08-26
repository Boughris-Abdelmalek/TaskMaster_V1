package api

import (
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/postgres"
)

// APIServer represents the API server structure
type APIServer struct {
	listenAddr string
	Store      postgres.Storage
}

// NewAPIServer creates a new instance of APIServer
func NewAPIServer(listenAddr string, Store postgres.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		Store:      Store,
	}
}
