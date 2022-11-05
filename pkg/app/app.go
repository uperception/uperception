package app

import (
	"github.com/leometzger/mmonitoring/pkg/storage"
	"github.com/leometzger/mmonitoring/pkg/storage/sqlstore"
)

type App struct {
	Config AppConfig
	store  storage.Store
}

type AppConfig struct{}

func NewApp() *App {
	return &App{
		Config: AppConfig{},
		store:  sqlstore.NewSQLStore(),
	}
}
