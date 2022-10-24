package store

import (
	"context"

	"github.com/HT0323/go_todo_app/entity"
)

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task
			(user_id, title, status, created, modified)
	VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, t.UserID, t.Title, t.Status,
		t.Created, t.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer, id entity.UserID,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
				id, user_id, title,
				status, created, modified
			FROM task
			WHERE user_id = ?;`
	if err := db.SelectContext(ctx, &tasks, sql, id); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) GetTask(
	ctx context.Context, db Queryer, id entity.UserID, taskId int,
) (entity.Task, error) {
	task := entity.Task{}
	sql := `SELECT
				id, user_id, title,
				status, created, modified
			FROM task
			WHERE user_id = ?
			AND id = ?;`
	if err := db.GetContext(ctx, &task, sql, id, taskId); err != nil {
		return entity.Task{}, err
	}
	return task, nil
}
