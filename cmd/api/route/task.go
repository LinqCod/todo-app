package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/linqcod/todo-app/internal/handler"
	repository "github.com/linqcod/todo-app/internal/repository/postgres"
	"github.com/linqcod/todo-app/internal/usecase"
	"github.com/linqcod/todo-app/pkg/database"
	"log/slog"
	"time"
)

func NewTaskRouter(logger *slog.Logger, timeout time.Duration, db *database.Postgres, validate *validator.Validate, group *gin.RouterGroup) {
	taskRepo := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo, timeout)
	taskHandler := handler.NewTaskHandler(logger, validate, taskUsecase)

	group.POST("/tasks", taskHandler.Create)
	group.DELETE("/tasks/:id", taskHandler.Delete)
	group.PATCH("/tasks/:id", taskHandler.Update)
	group.GET("/tasks/:id", taskHandler.GetById)
	group.GET("/tasks", taskHandler.GetFilteredWithPagination)
}
