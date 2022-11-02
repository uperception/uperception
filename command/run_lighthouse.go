package command

import (
	"log"

	"github.com/leometzger/mmonitoring-runner/collectors"
	"github.com/leometzger/mmonitoring-runner/storage"
	"github.com/spf13/cobra"
)

// Collect lighthouse metrics
func RunLighthouse(storage storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use: "run-lighthouse",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Start running lighthouse")
			lighthouse := collectors.NewLighthouseCollector(storage)

			log.Println("Created collector")
			err := lighthouse.Collect([]string{
				"https://google.com",
				"https://metzger.fot.br",
			})
			if err != nil {
				log.Fatal("error collecting lighthouse data", err)
			}

			log.Println("Finalized run")
		},
	}
}
