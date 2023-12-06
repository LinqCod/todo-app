package usecase

import (
	"context"
	"github.com/linqcod/todo-app/internal/domain"
	"time"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepo domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepo,
		contextTimeout: timeout,
	}
}

func (t taskUsecase) Create(ctx context.Context, task *domain.Task) error {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.Create(c, task)
}

func (t taskUsecase) Delete(ctx context.Context, id int64) (int64, error) {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.Delete(c, id)
}

func (t taskUsecase) Update(ctx context.Context, task *domain.UpdateTaskRequest) (*domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.Update(c, task)
}

func (t taskUsecase) GetById(ctx context.Context, id int64) (*domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.GetById(c, id)
}

func (t taskUsecase) GetFilteredWithPagination(ctx context.Context, offset, limit, date, isCompleted string) ([]*domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, t.contextTimeout)
	defer cancel()

	return t.taskRepository.GetFilteredWithPagination(c, offset, limit, date, isCompleted)
}
