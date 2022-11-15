package main

import (
	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/sql"
)

// Runs in a container environment that has
// chrome installed
// It must run in other process besides the API
func main() {
	sql.SetupModels(sql.Postgres)
	app := app.NewApp()
	app.CollectLighthouseData()
}
