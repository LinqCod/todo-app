package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/linqcod/todo-app/cmd/api/route"
	"github.com/linqcod/todo-app/internal/app/httpserver"
	"github.com/linqcod/todo-app/internal/config"
	"github.com/linqcod/todo-app/pkg/database"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg    *config.Config
	logger *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		cfg:    cfg,
		logger: log,
	}
}

func (a *App) Run() {
	const op = "app.Run"

	db, err := database.New(
		a.cfg.Postgres.Username,
		a.cfg.Postgres.Password,
		a.cfg.Postgres.Host,
		a.cfg.Postgres.DBName,
		a.cfg.Postgres.Port,
	)
	if err != nil {
		panic(err)
	}

	validate := validator.New()
	router := gin.Default()

	route.Setup(a.logger, a.cfg.HTTPServer.Timeout, db, validate, router)

	server := httpserver.New(a.logger, a.cfg.HTTPServer, router)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		server.MustRun()
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.HTTPServer.Timeout)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		a.logger.Error("error while stopping transport server: ", err)
		os.Exit(1)
	}

	if err := db.Close(); err != nil {
		a.logger.Error("error while closing postgres db connection: ", err)
		os.Exit(1)
	}

	a.logger.Info("app gracefully stopped!")
}
