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
		Use:  "run-lighthouse [url]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Start running lighthouse")
			lighthouse := collectors.NewLighthouseCollector(storage)

			log.Println("Created collector")
			err := lighthouse.Collect(args[0])
			if err != nil {
				log.Fatal("error collecting lighthouse data", err)
			}

			log.Println("Finalized run")
		},
	}
}
