package repository

import (
	"test/todo-app"

	"github.com/jmoiron/sqlx"
)

type Authoriztaion interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
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

type Repository struct { //Здесь из за того что имена и тип одинаковый можем не дублировать пхду?
	Authoriztaion
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authoriztaion: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
