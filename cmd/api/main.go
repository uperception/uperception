package main

import (
	"log"

	"github.com/leometzger/mmonitoring/pkg/api"
	"github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/db"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	db.SetupModels(db.Postgres)
	api.NewApi(config).Run()
}
