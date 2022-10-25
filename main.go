package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/leometzger/mmonitoring-runner/command"
	"github.com/leometzger/mmonitoring-runner/storage"
	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	storage := storage.NewAwsStorage(client, "mmontitoring")

	root.AddCommand(command.RunLighthouse(storage))
	root.Execute()
}
