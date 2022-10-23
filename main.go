package main

import (
	"github.com/leometzger/mmonitoring-runner/command"
	"github.com/leometzger/mmonitoring-runner/storage"
	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	storage := storage.NewAwsStorage()
	root.AddCommand(command.RunLighthouse(storage))
	root.Execute()
}
