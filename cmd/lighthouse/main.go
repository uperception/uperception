package main

import (
	"log"

	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/sql"
)

// Runs in a container environment that has
// chrome installed
// It must run in other process besides the API
func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	sql.SetupModels(sql.Postgres)
	app := app.NewApp(config)
	app.CollectLighthouseData()
}
