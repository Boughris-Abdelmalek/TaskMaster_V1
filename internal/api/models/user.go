package models

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	UserType  string    `json:"user_type"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewAccount(first_name, last_name, password, email, user_name string) (*User, error) {
	return &User{
		FirstName: first_name,
		LastName:  last_name,
		Password:  password,
		UserName:  user_name,
		CreatedAt: time.Now().UTC(),
		Email:     email,
		UserType:  "regular",
	}, nil
}
