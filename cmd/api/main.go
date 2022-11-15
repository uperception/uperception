package main

import (
	"log"

	"github.com/leometzger/mmonitoring/pkg/api"
	"github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/sql"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	sql.SetupModels(sql.Postgres)
	api.NewApi(config).Run()
}
