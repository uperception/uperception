package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/leometzger/mmonitoring/pkg/collectors"
	mConfig "github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/queue"
	"github.com/leometzger/mmonitoring/pkg/storage"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	appConfig, err := mConfig.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("Unable to load config: " + err.Error())
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(appConfig.Region))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	s3Client := s3.NewFromConfig(cfg)
	sqsClient := sqs.NewFromConfig(cfg)
	queue := queue.NewAwsQueue(sqsClient, appConfig.Queue)
	storage := storage.NewAwsStorage(s3Client, appConfig.Bucket)
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

}
