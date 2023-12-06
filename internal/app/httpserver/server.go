package httpserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/todo-app/internal/config"
	"log/slog"
	"net/http"
)

type Server struct {
	log        *slog.Logger
	httpServer *http.Server
}

func New(log *slog.Logger, cfg config.HTTPServerConfig, router *gin.Engine) *Server {
	//TODO: logger

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &Server{
		log:        log,
		httpServer: server,
	}
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "httpserver.Run"

	s.log.Info("http server starting", slog.String("addr", s.httpServer.Addr))

	if err := s.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	const op = "httpserver.Stop"

	s.log.Info("stopping http server")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("http server stopped")

	return nil
}
