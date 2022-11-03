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
	"github.com/leometzger/mmonitoring-runner/utils"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	appConfig, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load config", err)
	}

	root := cobra.Command{}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(appConfig.Region))
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	sqsClient := sqs.NewFromConfig(cfg)
	queue := queue.NewAwsQueue(sqsClient, appConfig.Queue)
	storage := storage.NewAwsStorage(s3Client, appConfig.Bucket)

	root.AddCommand(command.RunLighthouse(queue, storage))
	root.Execute()
}
