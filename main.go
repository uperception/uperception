package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/leometzger/mmonitoring-runner/command"
	"github.com/leometzger/mmonitoring-runner/queue"
	"github.com/leometzger/mmonitoring-runner/storage"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	root := cobra.Command{}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	sqsClient := sqs.NewFromConfig(cfg)
	queue := queue.NewAwsQueue(sqsClient, "")
	storage := storage.NewAwsStorage(s3Client, "mmonitoring")

	root.AddCommand(command.RunLighthouse(queue, storage))
	root.Execute()
}
