package app

import (
	"github.com/leometzger/mmonitoring/pkg/sql"
)

type App struct {
	Config AppConfig
	// queue                 queue.Queue
	// storage               storage.Storage
	// userStore             sql.UserStore
	projectStore          sql.ProjectStore
	organizationStore     sql.OrganizationStore
	sessionsStore         sql.SessionStore
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
