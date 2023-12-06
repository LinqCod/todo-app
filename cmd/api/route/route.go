package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/linqcod/todo-app/pkg/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"time"
)

func Setup(logger *slog.Logger, timeout time.Duration, db *database.Postgres, validate *validator.Validate, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")

	publicRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	NewTaskRouter(logger, timeout, db, validate, publicRouter)
}
