package service

import (
	"backend-mytodo/domain"
	"backend-mytodo/repository"
	"backend-mytodo/shareddomain"
)

type TodoServiceImpl struct {
	repository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{repository}
}

func (service *TodoServiceImpl) Create(request shareddomain.RequestCreate) error {
	todo := domain.Todo{
		Title: request.Title,
	}

	if err := service.repository.Create(todo); err != nil {
		return err
	}

	return nil
}

func (service *TodoServiceImpl) FindAll() ([]domain.Todo, error) {
	todos, err := service.repository.FindAll()
	if err != nil {
		return todos, err
	}
	return todos, nil
}