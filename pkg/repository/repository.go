package repository

import "github.com/jmoiron/sqlx"

type Authoriztaion interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct { //Здесь из за того что имена и тип одинаковый можем не дублировать пхду?
	Authoriztaion
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
