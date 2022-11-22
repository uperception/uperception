package app_test

import (
	"testing"

	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/testlib"
	"github.com/stretchr/testify/assert"
)

func TestLighthouseEndpointsBasicOperations(t *testing.T) {
	db.SetupModels(db.SQLite)
	defer testlib.ResetDatabase()
	app := app.NewApp(&config.Config{})

	_, err := app.CreateProject(models.CreateProjectInput{
		Name:        "Testing",
		Description: "Testing Desc",
	})
	assert.NoError(t, err)

	// Create
	endpoint, err := app.CreateLighthouseEndpoint("1", models.LighthouseEndpointInput{
		Url:    "https://google.com",
		ID:     0,
		Header: "",
	})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), endpoint.ID)

	// Find
	endpoint, err = app.FindLighthouseEndpoint("1")
	assert.NoError(t, err)
	assert.Equal(t, endpoint.Url, "https://google.com")

	// Update
	endpoint, err = app.UpdateLighthouseEndpoint(
		"1",
		models.LighthouseEndpointInput{
			Url:    "https://metzger.fot.br",
			ID:     1,
			Header: "Basic authtoken",
		})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), endpoint.ID)
	assert.Equal(t, "https://metzger.fot.br", endpoint.Url)
	assert.Equal(t, "Basic authtoken", endpoint.Header)

	// Delete
	err = app.DeleteLighthouseEndpoint("1")
	assert.NoError(t, err)

	_, err = app.FindLighthouseEndpoint("1")
	assert.Error(t, err)
}

func TestLighthouseBatchInsert(t *testing.T) {}
