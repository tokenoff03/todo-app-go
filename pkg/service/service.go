package service

import "test/todo-app/pkg/repository"

type Authoriztaion interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct { //Здесь из за того что имена и тип одинаковый можем не дублировать пхду?
	Authoriztaion
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
