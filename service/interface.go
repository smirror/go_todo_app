package service

import (
	"context"
	"todo_app/entity"
	"todo_app/store"
)

// go:generate go run github.com/matryer/moq -out mock_store_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
