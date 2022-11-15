package main

import (
	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/sql"

	"github.com/rs/zerolog"
)

// Runs in a container environment that has
// chrome installed
// It must run in other process besides the API
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	sql.SetupModels(sql.Postgres)
	app := app.NewApp()
	app.CollectLighthouseData()
}
