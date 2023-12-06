package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/linqcod/todo-app/internal/domain"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockTaskUsecase struct {
	t *testing.T
}

var (
	tasks = []domain.Task{
		{Id: 1, Title: "Task 1", Description: "Its task 1", Date: "2023-12-08", IsCompleted: false},
		{Id: 2, Title: "Task 2", Description: "Its task 2", Date: "2023-12-05", IsCompleted: true},
		{Id: 3, Title: "Task 3", Description: "Its task 3", Date: "2023-12-02", IsCompleted: true},
		{Id: 4, Title: "Task 4", Description: "Its task 4", Date: "2024-01-23", IsCompleted: false},
		{Id: 5, Title: "Task 5", Description: "Its task 5", Date: "2023-12-24", IsCompleted: true},
		{Id: 6, Title: "Task 6", Description: "Its task 6", Date: "2023-12-27", IsCompleted: false},
		{Id: 7, Title: "Task 7", Description: "Its task 7", Date: "2023-12-24", IsCompleted: false},
		{Id: 8, Title: "Task 8", Description: "Its task 8", Date: "2023-12-05", IsCompleted: true},
		{Id: 9, Title: "Task 9", Description: "Its task 9", Date: "2023-12-05", IsCompleted: false},
	}
	notValidTask = domain.Task{Id: 1, Title: "Task 1", Description: "Its task 1", Date: "waefrgrgrgr", IsCompleted: false}
	wantError    bool
)

func NewMockTaskUsecase(t *testing.T) *MockTaskUsecase {
	return &MockTaskUsecase{t: t}
}

func (t MockTaskUsecase) Create(ctx context.Context, task *domain.Task) error {
	if task.Id <= 0 || task.Title == "" || task.Description == "" || task.Date == "" {
		return errors.New("invalid task")
	}

	return nil
}

func (t MockTaskUsecase) Delete(ctx context.Context, id int64) (int64, error) {
	if int64(len(tasks)) < id {
		return -1, errors.New("task with given id not found")
	}

	return id, nil
}

func (t MockTaskUsecase) Update(ctx context.Context, task *domain.UpdateTaskRequest) (*domain.Task, error) {
	if int64(len(tasks)) < task.Id {
		return nil, errors.New("task with given id not found")
	}

	return &tasks[task.Id-1], nil
}

func (t MockTaskUsecase) GetById(ctx context.Context, id int64) (*domain.Task, error) {
	if int64(len(tasks)) < id {
		return nil, errors.New("task with given id not found")
	}

	return &tasks[id-1], nil
}

func (t MockTaskUsecase) GetFilteredWithPagination(ctx context.Context, offset, limit, date, isCompleted string) ([]*domain.Task, error) {
	if wantError {
		return nil, errors.New("error found")
	}

	res := make([]*domain.Task, len(tasks))
	for _, v := range tasks {
		res = append(res, &v)
	}
	return res, nil
}

func NewTestHandler(t *testing.T) *TaskHandler {
	t.Helper()

	gin.SetMode(gin.TestMode)

	logger := &slog.Logger{}
	validate := validator.New()

	mock := NewMockTaskUsecase(t)
	h := NewTaskHandler(logger, validate, mock)

	return h
}

func NewTestRecordWriter() (*httptest.ResponseRecorder, *gin.Context) {
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)

	return writer, ctx
}

func TestTaskCreateHandler(t *testing.T) {
	h := NewTestHandler(t)

	// EXPECT CREATED
	// should return status 201 Created
	t.Run("EXPECT CREATED", func(t *testing.T) {
		writer, ctx := NewTestRecordWriter()

		taskJSON, err := json.Marshal(&tasks[0])
		assert.NoError(t, err)

		ctx.Request, err = http.NewRequest("POST", "/", bytes.NewBuffer(taskJSON))
		assert.NoError(t, err)

		ctx.Request.Header.Add("content-type", "application/json")

		h.Create(ctx)

		assert.Equal(t, http.StatusCreated, writer.Code)
	})
}

func TestTaskGetHandler(t *testing.T) {
	handler := NewTestHandler(t)

	// EXPECT SUCCESS
	// should return status 200 ok
	t.Run("EXPECT SUCCESS", func(t *testing.T) {
		writer, ctx := NewTestRecordWriter()
		ctx.Params = gin.Params{
			{Key: "id", Value: "1"},
		}

		handler.GetById(ctx)

		assert.Equal(t, http.StatusOK, writer.Code)
	})
}
