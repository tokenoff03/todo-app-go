package service

import (
	"test/todo-app"
	"test/todo-app/pkg/repository"
)

type Authoriztaion interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId int, id int, input todo.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct { //Здесь из за того что имена и тип одинаковый можем не дублировать пхду?
	Authoriztaion
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authoriztaion: NewAuthService(repos.Authoriztaion),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
