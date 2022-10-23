package handler

import "github.com/HT0323/go_todo_app/entity"

type Task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}
