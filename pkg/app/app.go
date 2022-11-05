package app

import (
	"github.com/leometzger/mmonitoring/pkg/sql"
)

type App struct {
	Config AppConfig
	store  sql.Store
}

type AppConfig struct{}

func NewApp() *App {
	return &App{
		Config: AppConfig{},
		store:  sql.NewSQLStore(),
	}
}
