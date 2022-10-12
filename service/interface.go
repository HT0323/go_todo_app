package service

import (
	"context"

	"github.com/HT0323/go_todo_app/entity"
	"github.com/HT0323/go_todo_app/store"
)

type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
