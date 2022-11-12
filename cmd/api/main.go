package main

import (
	"github.com/leometzger/mmonitoring/pkg/api"
	"github.com/leometzger/mmonitoring/pkg/sql"
)

func main() {
	sql.SetupModels(sql.Postgres)
	api.NewApi().Run()
}
