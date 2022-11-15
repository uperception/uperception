package app

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/leometzger/mmonitoring/pkg/collectors"
	mConfig "github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/queue"
	"github.com/leometzger/mmonitoring/pkg/sql"
	"github.com/leometzger/mmonitoring/pkg/storage"
)

type App struct {
	// userStore             sql.UserStore
	config                *mConfig.Config
	queue                 queue.Queue
	storage               storage.Storage
	lighthouseCollector   collectors.Collector
	projectStore          sql.ProjectStore
	organizationStore     sql.OrganizationStore
	sessionsStore         sql.SessionStore
	lighthouseResultStore sql.LighthouseResultStore
	lighthouseConfigStore sql.LighthouseConfigStore
}

func NewApp(appConfig *mConfig.Config) *App {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(appConfig.Region))
	if err != nil {
		log.Fatal(err)
	}

	sqsClient := sqs.NewFromConfig(cfg)
	q := queue.NewAwsQueue(sqsClient, appConfig.Queue, appConfig.QueueUrl)

	s3Client := s3.NewFromConfig(cfg)
	storage := storage.NewAwsStorage(s3Client, appConfig.Bucket)
	lighthouse := collectors.NewLighthouseCollector(storage, sql.NewLighthouseResultStore())

	return &App{
		config:                appConfig,
		queue:                 q,
		lighthouseCollector:   lighthouse,
		projectStore:          sql.NewProjectStore(),
		organizationStore:     sql.NewOrganizationStore(),
		sessionsStore:         sql.NewSessionStore(),
		lighthouseResultStore: sql.NewLighthouseResultStore(),
		lighthouseConfigStore: sql.NewLighthouseConfigStore(),
	}
}
