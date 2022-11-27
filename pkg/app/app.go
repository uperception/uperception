package app

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v12"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/leometzger/mmonitoring/pkg/collectors"
	appConfig "github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/queue"
	"github.com/leometzger/mmonitoring/pkg/storage"
)

type App struct {
	// userStore             sql.UserStore
	config                   *appConfig.Config
	keycloakClient           *gocloak.GoCloak
	keycloakAdminToken       *gocloak.JWT
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

func NewApp(appConfig *appConfig.Config) *App {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(appConfig.Region))
	if err != nil {
		log.Fatal(err)
	}

	sqsClient := sqs.NewFromConfig(cfg)
	q := queue.NewAwsQueue(sqsClient, appConfig.Queue, appConfig.QueueUrl)

	s3Client := s3.NewFromConfig(cfg)
	storage := storage.NewAwsStorage(s3Client, appConfig.Bucket)
	lighthouse := collectors.NewLighthouseCollector(storage, db.NewLighthouseResultStore())

	client := gocloak.NewClient(appConfig.KeycloakUrl)

	return &App{
		config:                   appConfig,
		queue:                    q,
		lighthouseCollector:      lighthouse,
		keycloakClient:           client,
		storage:                  storage,
		projectStore:             db.NewProjectStore(),
		organizationStore:        db.NewOrganizationStore(),
		sessionsStore:            db.NewSessionStore(),
		lighthouseResultStore:    db.NewLighthouseResultStore(),
		lighthouseConfigStore:    db.NewLighthouseConfigStore(),
		lighthouseEndpointsStore: db.NewLighthouseEndpointsStore(),
	}
}

func (a *App) refreshKeycloakToken() {
	ctx := context.Background()
	token, err := a.keycloakClient.LoginClient(ctx, a.config.KeycloakClient, a.config.KeycloakSecret, a.config.KeycloakRealm)
	if err != nil {
		return
	}
	a.keycloakAdminToken = token
}
