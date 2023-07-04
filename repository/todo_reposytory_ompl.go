package repository

import (
	"backend-mytodo/domain"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{db}
}

func (repo *TodoRepositoryImpl) Create(todo domain.Todo) error {
	if err := repo.db.Create(&todo).Error; err != nil {
		return err
	}

	return nil
}
func (repo *TodoRepositoryImpl) FindAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	if err := repo.db.Create(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}