package main

import (
	"github.com/leometzger/mmonitoring/pkg/api"
	"github.com/leometzger/mmonitoring/pkg/db"
)

func main() {
	db.SetupModels(db.Postgres)
	api.NewApi().Run()
}
