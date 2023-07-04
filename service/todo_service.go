package service

import "backend-mytodo/shareddomain"

type TodoService interface {
	Create(request shareddomain.RequestCreate) error
}