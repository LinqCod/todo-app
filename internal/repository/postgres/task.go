package repository

import (
	"context"
	"fmt"
	"github.com/linqcod/todo-app/internal/domain"
	"github.com/linqcod/todo-app/pkg/database"
	"log/slog"
	"strings"
)

type taskRepository struct {
	postgres *database.Postgres
}

func NewTaskRepository(db *database.Postgres) domain.TaskRepository {
	return &taskRepository{
		postgres: db,
	}
}

func (t taskRepository) Create(ctx context.Context, task *domain.Task) error {
	if err := t.postgres.DB.QueryRowContext(
		ctx,
		`
			INSERT INTO tasks (title, description, assigned_date, is_completed) 
			VALUES ($1, $2, $3, $4);
		`,
		task.Title,
		task.Description,
		task.Date,
		task.IsCompleted,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (t taskRepository) Delete(ctx context.Context, id int64) (int64, error) {
	if err := t.postgres.DB.QueryRowContext(
		ctx,
		` DELETE FROM tasks WHERE id=$1 RETURNING id;`,
		id,
	).Scan(
		&id,
	); err != nil {
		return -1, err
	}

	return id, nil
}

func (t taskRepository) Update(ctx context.Context, task *domain.UpdateTaskRequest) (*domain.Task, error) {
	var updatedTask domain.Task

	if err := t.postgres.DB.QueryRowContext(
		ctx,
		`
			UPDATE tasks
			SET title=coalesce($1, title), 
			    description=coalesce($2, description), 
			    assigned_date=coalesce($3, assigned_date), 
			    is_completed=coalesce($4, is_completed)
			WHERE id=$5 RETURNING id, title, description, assigned_date, is_completed;
		`,
		task.Title,
		task.Description,
		task.Date,
		task.IsCompleted,
		task.Id,
	).Scan(
		&updatedTask.Id,
		&updatedTask.Title,
		&updatedTask.Description,
		&updatedTask.Date,
		&updatedTask.IsCompleted,
	); err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

func (t taskRepository) GetById(ctx context.Context, id int64) (*domain.Task, error) {
	var task domain.Task

	if err := t.postgres.DB.QueryRowContext(
		ctx,
		`
			SELECT id, title, description, assigned_date, is_completed
 			FROM tasks 
 			WHERE id=$1;
		`,
		id,
	).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Date,
		&task.IsCompleted,
	); err != nil {
		return nil, err
	}

	return &task, nil
}

func (t taskRepository) GetFilteredWithPagination(
	ctx context.Context,
	offset,
	limit,
	date,
	isCompleted string,
) ([]*domain.Task, error) {
	var tasks []*domain.Task

	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, title, description, assigned_date, is_completed FROM tasks")

	if date != "" || isCompleted != "" {
		queryBuilder.WriteString(" WHERE ")
		conditions := make([]string, 0)
		if date != "" {
			conditions = append(conditions, fmt.Sprintf("assigned_date='%s'", date))
			//TODO: work with date in postgres
		}
		if isCompleted != "" {
			conditions = append(conditions, fmt.Sprintf("is_completed=%s", isCompleted))
		}

		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}

	if limit != "" {
		queryBuilder.WriteString(fmt.Sprintf(" LIMIT %s", limit))
	}
	if offset != "" {
		queryBuilder.WriteString(fmt.Sprintf(" OFFSET %s", offset))
	}

	slog.Debug("Final query", queryBuilder.String())

	rows, err := t.postgres.DB.QueryContext(
		ctx,
		queryBuilder.String(),
	)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task

		if err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Date,
			&task.IsCompleted,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}
