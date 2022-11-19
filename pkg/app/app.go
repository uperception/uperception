package app

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/leometzger/mmonitoring/pkg/collectors"
	mConfig "github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/queue"
	"github.com/leometzger/mmonitoring/pkg/storage"
)

type App struct {
	// userStore             sql.UserStore
	config                   *mConfig.Config
	queue                    queue.Queue
	storage                  storage.Storage
	lighthouseCollector      collectors.Collector
	projectStore             db.ProjectStore
	organizationStore        db.OrganizationStore
	sessionsStore            db.SessionStore
	lighthouseResultStore    db.LighthouseResultStore
	lighthouseConfigStore    db.LighthouseConfigStore
	lighthouseEndpointsStore db.LighthouseEndpointsStore
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
	lighthouse := collectors.NewLighthouseCollector(storage, db.NewLighthouseResultStore())

	return &App{
		config:                   appConfig,
		queue:                    q,
		lighthouseCollector:      lighthouse,
		projectStore:             db.NewProjectStore(),
		organizationStore:        db.NewOrganizationStore(),
		sessionsStore:            db.NewSessionStore(),
		lighthouseResultStore:    db.NewLighthouseResultStore(),
		lighthouseConfigStore:    db.NewLighthouseConfigStore(),
		lighthouseEndpointsStore: db.NewLighthouseEndpointsStore(),
	}
}
