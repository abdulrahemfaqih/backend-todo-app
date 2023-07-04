package service

import "backend-mytodo/shareddomain"
import "backend-mytodo/domain"

type TodoService interface {
	Create(request shareddomain.RequestCreate) error
	FindAll() ([]domain.Todo, error)
}