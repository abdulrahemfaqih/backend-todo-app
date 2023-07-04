package repository

import "backend-mytodo/domain"

type TodoRepository interface {
	Create(todo domain.Todo) error
}