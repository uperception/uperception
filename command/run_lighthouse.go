package command

import (
	"log"
	"os/exec"

	"github.com/leometzger/mmonitoring-runner/storage"
	"github.com/spf13/cobra"
)

func RunLighthouse(storage storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use:  "run-lighthouse [domain]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			domain := args[0]

			err := exec.Command(
				"lighthouse",
				domain,
				"--output-path=./reports/"+"report"+".json",
				"--gather-mode",
				"--save-assets",
				"--output='json,html'",
				"--chrome-flags='--headless'",
			).Run()

			if err != nil {
				log.Fatal("Error running lighthouse command", err)
			}

			err = storage.SaveLighthouseResult(domain)
			if err != nil {
				log.Fatal("Error savind lighthouse result")
			}
		},
	}
}
