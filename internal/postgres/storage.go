package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/types"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbUsername = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbHost     = "DB_HOST"
	dbSchema   = "DB_NAME"
)

var (
	username = os.Getenv(dbUsername)
	password = os.Getenv(dbPassword)
	host     = os.Getenv(dbHost)
	schema   = os.Getenv(dbSchema)
)

type Storage interface {
	GetTodos() ([]types.Todo, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgres() (*PostgresStore, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, password, schema)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) GetTodos() ([]types.Todo, error) {
	var todos []types.Todo

	rows, err := s.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo types.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
