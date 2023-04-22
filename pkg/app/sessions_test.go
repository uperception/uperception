package app_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/testlib"
	"github.com/stretchr/testify/assert"
)

func TestStartSession(t *testing.T) {
	db.SetupModels(db.SQLite)
	defer testlib.ResetDatabase()
	app := app.NewApp(&config.Config{})
	project, _ := app.CreateProject(models.CreateProjectInput{
		Name:        "Testing Project",
		Description: "Testing Description",
	})

	output, err := app.StartSession(models.StartSessionInput{
		Token: project.Token,
		Path:  "/my-application-path",
	})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), output.ID)
}

func TestPublishEvents(t *testing.T) {
	db.SetupModels(db.SQLite)
	defer testlib.ResetDatabase()
	app := app.NewApp(&config.Config{})
	project, _ := app.CreateProject(models.CreateProjectInput{
		Name:        "Testing Project",
		Description: "Testing Description",
	})
	output, _ := app.StartSession(models.StartSessionInput{
		Token: project.Token,
		Path:  "/my-application-path",
	})
	data := []map[string]interface{}{
		{
			"type":      4,
			"timestamp": 1682186201869,
			"data":      "DataEvents",
		},
		{
			"type":      3,
			"timestamp": 1682186201870,
			"data":      "Event2",
		},
	}

	err := app.PublishEvents(models.PublishEventsInput{
		ID:         fmt.Sprintf("%d", output.ID),
		Path:       "/my-application-path",
		Timestamp:  uint(time.Now().Unix()),
		EventsData: data,
	})

	assert.NoError(t, err)
}

func TestFinishSession(t *testing.T) {
	db.SetupModels(db.SQLite)
	defer testlib.ResetDatabase()
}

func TestPublishOnFinishedSession(t *testing.T) {
	db.SetupModels(db.SQLite)
	defer testlib.ResetDatabase()
}
