package command

import (
	"github.com/leometzger/mmonitoring-runner/collectors"
	"github.com/leometzger/mmonitoring-runner/queue"
	"github.com/leometzger/mmonitoring-runner/storage"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Collect lighthouse metrics
func RunLighthouse(queue queue.Queue, storage storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use: "run-lighthouse",
		Run: func(cmd *cobra.Command, args []string) {
			lighthouse := collectors.NewLighthouseCollector(storage)

			task, err := queue.GetTask()
			if err != nil {
				log.Fatal().Msg("Getting Task: " + err.Error())
			}

			if len(task.Project.Urls) == 0 {
				return
			}

			err = lighthouse.Collect(task.Project.Urls)
			if err != nil {
				log.Fatal().Msg("Collecting Lighthouse Data:" + err.Error())
			}
		},
	}
}
