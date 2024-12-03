package service

import "todo-project/pkg/repository"

type Authorization interface {
}

type TodoLists interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoLists
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
