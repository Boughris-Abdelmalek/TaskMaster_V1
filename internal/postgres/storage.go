package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Boughris-Abdelmalek/TaskMaster_V1/internal/api"
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
	GetTodos() ([]api.Todo, error)
	GetTodoByID(int) (*api.Todo, error)
	CreateTodo(*api.Todo) error
	DeleteTodo(int) error
	UpdateTodo(int, *api.Todo) (*api.Todo, error)
}

type PostgresStore struct {
	db *sql.DB
}

// NewPostgres set up the connection to with the db
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

// GetTodos is a method that performs a query to get all the todos
func (s *PostgresStore) GetTodos() ([]api.Todo, error) {
	var todos []api.Todo

	rows, err := s.db.Query("SELECT * FROM todos ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo api.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// CreateTodo is a method that performs a query to create new todo
func (s *PostgresStore) CreateTodo(todo *api.Todo) error {
	query := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id`

	_, err := s.db.Query(query, todo.Title, todo.Description, todo.Completed)
	if err != nil {
		return err
	}

	return nil
}

// GetTodoByID is a method that performs a query to get one todo by ID
func (s *PostgresStore) GetTodoByID(id int) (*api.Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoTodos(rows)
	}

	return nil, fmt.Errorf("account with number [%d] not found", id)
}

// DeleteTodo is a method that performs a query to delete todos
func (s *PostgresStore) DeleteTodo(id int) error {
	_, err := s.db.Query("DELETE FROM todos WHERE id = $1", id)
	return err
}

// UpdateTodo is a method that performs a query to update todos
func (s *PostgresStore) UpdateTodo(id int, todoUpdate *api.Todo) (*api.Todo, error) {
	// Retrieve the existing todo from the database
	row, err := s.db.Query("SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		todo, err := scanIntoTodos(row)
		if err != nil {
			return nil, err
		}

		// Update the todo fields with the new values
		todo.Title = todoUpdate.Title
		todo.Description = todoUpdate.Description
		todo.Completed = todoUpdate.Completed

		// Update the todo in the database
		_, execErr := s.db.Exec("UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4",
			todo.Title, todo.Description, todo.Completed, id)
		if execErr != nil {
			return nil, execErr
		}

		return todo, nil
	}

	return nil, fmt.Errorf("todo with ID %d not found", id)
}

func scanIntoTodos(rows *sql.Rows) (*api.Todo, error) {
	todo := new(api.Todo)
	err := rows.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
	)
	return todo, err
}
