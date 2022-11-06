package app

import (
	"github.com/leometzger/mmonitoring/pkg/queue"
	"github.com/leometzger/mmonitoring/pkg/sql"
	"github.com/leometzger/mmonitoring/pkg/storage"
)

type App struct {
	Config                AppConfig
	queue                 queue.Queue
	storage               storage.Storage
	projectStore          sql.ProjectStore
	organizationStore     sql.OrganizationStore
	sessionsStore         sql.SessionStore
	userStore             sql.UserStore
	lighthouseResultStore sql.LighthouseResultStore
	lighthouseConfigStore sql.LighthouseConfigStore
}

type AppConfig struct{}

func NewApp() *App {
	return &App{
		Config:                AppConfig{},
		projectStore:          sql.NewProjectStore(),
		organizationStore:     sql.NewOrganizationStore(),
		sessionsStore:         sql.NewSessionStore(),
		lighthouseResultStore: sql.NewLighthouseResultStore(),
		lighthouseConfigStore: sql.NewLighthouseConfigStore(),
	}
}
