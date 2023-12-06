package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/linqcod/todo-app/internal/domain"
	"github.com/linqcod/todo-app/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	logger      *slog.Logger
	validate    *validator.Validate
	TaskUsecase domain.TaskUsecase
}

func NewTaskHandler(logger *slog.Logger, validate *validator.Validate, taskUsecase domain.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		logger:      logger,
		validate:    validate,
		TaskUsecase: taskUsecase,
	}
}

// Create godoc
//
//	@Summary		create task
//	@Description	create task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		domain.Task	true	"Create task"
//	@Success		201		{object}	domain.SuccessResponse	"Task created successfully"
//	@Failure		400		{object}	domain.ErrorResponse	"error bad request data"
//	@Failure		500		{object}	domain.ErrorResponse	"error while creating task"
//	@Router			/tasks [post]
func (h *TaskHandler) Create(c *gin.Context) {
	var task domain.Task

	if err := c.ShouldBind(&task); err != nil {
		h.logger.Error("error while binding request body", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.validate.Struct(task); err != nil {
		h.logger.Error("error while validating task body", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err := utils.IsTaskDateStringValid(task.Date)
	if err != nil {
		h.logger.Error("error while validating task date", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.TaskUsecase.Create(c, &task); err != nil {
		h.logger.Error("error while creating task", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

// Delete godoc
//
//	@Summary		delete task
//	@Description	delete task
//	@Tags			tasks
//	@Produce		json
//	@Param			id	path		int     true	"Task id"
//	@Success		200		{object}	domain.SuccessResponse	"Task deleted successfully"
//	@Failure		400		{object}	domain.ErrorResponse	"error bad request data"
//	@Failure		500		{object}	domain.ErrorResponse	"error while deleting task"
//	@Router			/tasks/{id} [delete]
func (h *TaskHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("error while getting task id", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if _, err := h.TaskUsecase.Delete(c, int64(id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.logger.Error("id not found", err)
			c.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: fmt.Sprintf("task with id: %d not found", id),
			})
		} else {
			h.logger.Error("error while deleting task", err)
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
				Message: err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: fmt.Sprintf("task with id: %d deleted successfully", id),
	})
}

// Update godoc
//
//	@Summary		update task
//	@Description	update task
//	@Tags			tasks
//	@Produce		json
//	@Param			id	path		int     true	"Task id"
//	@Param			task	body		domain.UpdateTaskRequest	true	"Update task"
//	@Success		200		{object}	domain.SuccessWithDataResponse	"Task updated successfully"
//	@Failure		400		{object}	domain.ErrorResponse	"error bad request data"
//	@Failure		500		{object}	domain.ErrorResponse	"error while updating task"
//	@Router			/tasks/{id} [patch]
func (h *TaskHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("error while getting task id", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var task domain.UpdateTaskRequest

	if err := c.ShouldBind(&task); err != nil {
		h.logger.Error("error while binding request body", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	task.Id = int64(id)

	if task.Date != nil {
		err = utils.IsTaskDateStringValid(*task.Date)
		if err != nil && *task.Date != "" {
			h.logger.Error("error while validating task date", err)
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
	}

	task.Id = int64(id)

	updatedTask, err := h.TaskUsecase.Update(c, &task)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.logger.Error("id not found", err)
			c.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: fmt.Sprintf("task with id: %d not found", id),
			})
		} else {
			h.logger.Error("error while updating task", err)
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
				Message: err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, domain.SuccessWithDataResponse{
		Message: "Task updated successfully",
		Data:    updatedTask,
	})
}

// GetById godoc
//
//	@Summary		get task by id
//	@Description	get task by id
//	@Tags			tasks
//	@Produce		json
//	@Param			id	path		int     true	"Task id"
//	@Success		200		{object}	domain.SuccessWithDataResponse	"Task got successfully"
//	@Failure		400		{object}	domain.ErrorResponse	"error bad request data"
//	@Failure		500		{object}	domain.ErrorResponse	"error while getting task by id"
//	@Router			/tasks/{id} [get]
func (h *TaskHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("error while getting task id", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	task, err := h.TaskUsecase.GetById(c, int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.logger.Error("id not found", err)
			c.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: fmt.Sprintf("task with id: %d not found", id),
			})
		} else {
			h.logger.Error("error while getting task by id", err)
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
				Message: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, domain.SuccessWithDataResponse{
		Message: "Task got successfully",
		Data:    task,
	})
}

// GetFilteredWithPagination godoc
//
//	@Summary		get tasks with filters and pagination
//	@Description	get tasks with filters and pagination
//	@Tags			tasks
//	@Produce		json
//	@Param			offset	query		int  false	"tasks pagination offset"
//	@Param			limit	query		int  false	"tasks pagination limit"
//	@Param			date	query		string  false	"task assigned date"
//	@Param			isCompleted	query		bool  false	"status of task"
//	@Success		200		{object}	domain.SuccessWithDataResponse	"Task updated successfully"
//	@Failure		400		{object}	domain.ErrorResponse	"error bad request data"
//	@Failure		500		{object}	domain.ErrorResponse	"error while getting filtered tasks"
//	@Router			/tasks [get]
func (h *TaskHandler) GetFilteredWithPagination(c *gin.Context) {
	offset := c.Query("offset")

	offsetValue, err := strconv.Atoi(offset)
	if err != nil && offset != "" {
		h.logger.Error("error while parsing offset", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	limit := c.Query("limit")

	limitValue, err := strconv.Atoi(limit)
	if err != nil && limit != "" {
		h.logger.Error("error while parsing limit", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if limitValue < 0 || offsetValue < 0 {
		h.logger.Error("error: limit and offset values can't be less then zero")
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "error: limit and offset values can't be less then zero",
		})
		return
	}

	date := c.Query("date")

	_, err = utils.GetFormattedDateFromString(date)
	if err != nil && date != "" {
		h.logger.Error("error while validating task date", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	isCompleted := c.Query("isCompleted")

	tasks, err := h.TaskUsecase.GetFilteredWithPagination(c, offset, limit, date, isCompleted)
	if err != nil {
		h.logger.Error("error while getting filtered tasks", err)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessWithDataResponse{
		Message: "Users got successfully",
		Data:    tasks,
	})
}
