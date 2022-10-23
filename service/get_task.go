package service

import (
	"context"
	"fmt"

	"github.com/HT0323/go_todo_app/auth"
	"github.com/HT0323/go_todo_app/entity"
	"github.com/HT0323/go_todo_app/store"
)

type GetTask struct {
	DB   store.Queryer
	Repo TaskGeter
}

func (g *GetTask) GetTask(ctx context.Context, taskID int) (entity.Task, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return entity.Task{}, fmt.Errorf("user_id not found")
	}
	ts, err := g.Repo.GetTask(ctx, g.DB, id, taskID)
	if err != nil {
		return entity.Task{}, fmt.Errorf("failed to task: %w", err)
	}
	return ts, nil
}
