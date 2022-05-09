package repository

import (
	"database/sql"
)

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

/* Конструктор БД
Так как работает с БД передаем объект sqlx.DB в качестве аргумента */
func NewRepository(*sql.DB) *Repository {
	return &Repository{}
}
