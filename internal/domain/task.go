package domain

import (
	"context"
)

type Task struct {
	Id          int64  `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Date        string `json:"assigned_date" validate:"required"`
	IsCompleted bool   `json:"is_completed"`
}

type UpdateTaskRequest struct {
	Id          int64   `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Date        *string `json:"assigned_date"`
	IsCompleted *bool   `json:"is_completed"`
}

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id int64) (int64, error)
	Update(ctx context.Context, task *UpdateTaskRequest) (*Task, error)
	GetById(ctx context.Context, id int64) (*Task, error)
	GetFilteredWithPagination(ctx context.Context, offset, limit, date, isCompleted string) ([]*Task, error)
}

type TaskUsecase interface {
	Create(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id int64) (int64, error)
	Update(ctx context.Context, task *UpdateTaskRequest) (*Task, error)
	GetById(ctx context.Context, id int64) (*Task, error)
	GetFilteredWithPagination(ctx context.Context, offset, limit, date, isCompleted string) ([]*Task, error)
}
