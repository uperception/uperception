package app_test

import (
	"strconv"
	"testing"

	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/testlib"
	"github.com/stretchr/testify/assert"
)

func TestProjectBasicOperations(t *testing.T) {
	db.SetupModels(db.SQLite)
	defer testlib.ResetDatabase()
	app := app.NewApp(&config.Config{})

	// Create
	project, err := app.CreateProject(models.CreateProjectInput{
		Name:        "Testing Project",
		Description: "Testing Description",
	})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), project.ID)

	// Find
	project, err = app.FindProject("1")
	assert.NoError(t, err)
	assert.Equal(t, project.Name, "Testing Project")

	// Update
	project, err = app.UpdateProject("1", models.UpdateProjectInput{
		Name:        "Testing Project v2",
		Description: "Testing Description v2",
	})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), project.ID)
	assert.Equal(t, project.Name, "Testing Project v2")

	// Delete
	err = app.DeleteProject("1")
	assert.NoError(t, err)

	_, err = app.FindProject("1")
	assert.Error(t, err)
}

func TestUpdateProjectLighthouseConfig(t *testing.T) {
	db.SetupModels(db.SQLite)
	defer testlib.ResetDatabase()
	app := app.NewApp(&config.Config{})

	project, err := app.CreateProject(models.CreateProjectInput{
		Name:        "Testing Project With Lighthouse Config",
		Description: "Testing Description",
	})
	assert.NoError(t, err)
	projectID := strconv.FormatUint(uint64(project.ID), 10)

	config, err := app.UpdateLighthouseConfig(
		projectID,
		&models.UpdateLighthouseConfigInput{
			Enabled:     true,
			Periodicity: 1,
			Endpoints: []models.LighthouseEndpointInput{
				{Url: "https://google.com"},
				{Url: "https://private.com", Header: "Basic 1234"},
			},
		})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), config.ID)

	project, err = app.FindProject(projectID)
	assert.NoError(t, err)
	assert.Equal(t, config.ID, project.LighthouseConfig.ID)
	assert.Equal(t, 2, len(project.LighthouseConfig.Endpoints))

	_, err = app.UpdateLighthouseConfig(
		projectID,
		&models.UpdateLighthouseConfigInput{
			Enabled:     false,
			Periodicity: 1,
			Endpoints: []models.LighthouseEndpointInput{
				{ID: 1, Url: "https://google.com"},
				{ID: 2, Url: "https://facebook.com"},
				{Url: "https://private.com", Header: "Basic 1234"},
			},
		})
	assert.NoError(t, err)

	project, err = app.FindProject(projectID)
	assert.NoError(t, err)
	assert.Equal(t, false, project.LighthouseConfig.Enabled)
	assert.Equal(t, 3, len(project.LighthouseConfig.Endpoints))
}
