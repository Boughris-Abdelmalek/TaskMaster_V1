package models

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewTodo(title, description string) (*Todo, error) {
	return &Todo{
		Title:       title,
		Description: description,
		Completed:   false,
	}, nil
}

func NewTodoUpdate(title, description string, completed bool) (*Todo, error) {
	return &Todo{
		Title:       title,
		Description: description,
		Completed:   completed,
	}, nil
}
