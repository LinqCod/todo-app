package main

import (
	_ "github.com/linqcod/todo-app/docs"
	"github.com/linqcod/todo-app/internal/app"
	"github.com/linqcod/todo-app/internal/config"
	"github.com/linqcod/todo-app/pkg/logger/sl"
)

// @title TODO API
// @version 1.0
// @description todo service

// @contact.name linqcod
// @contact.email linqcod@yandex.ru

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	app.New(cfg, log).Run()
}
