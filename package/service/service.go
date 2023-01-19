package service

import (
	"github.com/olegvolkov91/go-todo-rest-tutorial"
	"github.com/olegvolkov91/go-todo-rest-tutorial/package/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r),
	}
}
