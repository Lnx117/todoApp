package service

import (
	"todo"
	"todo/pkg/repository"
)

//Принимаем данные с уровня запросов
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
	//В Authorization должен лежать объект содержащий метод CreateUser
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	// Во время создания цепочки объектов в мейне этот конструктор положит в Authorization уже реализацию этого интерфейса
	//а именно NewAuthService
	//И так со всем остальным
	return &Service{
		//В этом объекте должна быть функция которая получает на входе данные в виде todo.User
		//Все так и есть
		Authorization: NewAuthService(repos.Authorization),
	}
}
