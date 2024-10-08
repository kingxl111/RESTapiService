package service

import (
	todo "github.com/kingxl111/RESTapiService"
	"github.com/kingxl111/RESTapiService/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	// Передаем на следующий уровень
	return s.repo.Create(userId, list)
}