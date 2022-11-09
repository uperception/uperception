package app_test

import (
	"strconv"
	"testing"

	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/sql"
	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	sql.SetupModels(sql.SQLite)
	// defer testlib.ResetDatabase()
	app := app.NewApp()

	project, err := app.CreateProject(models.CreateProjectInput{
		Name:        "Testing Project",
		Description: "Testing Description",
	})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), project.ID)
}

// func TestUpdateProject(t *testing.T) {
// 	sql.SetupModels(sql.SQLite)
// 	defer testlib.ResetDatabase()
// 	app := NewApp()
//
// 	project, err := app.UpdateProject("1", models.UpdateProjectInput{
// 		Name:        "Testing Project v2",
// 		Description: "Testing Description v2",
// 	})
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, uint(1), project.ID)
// }

func TestUpdateLighthouseConfig(t *testing.T) {
	sql.SetupModels(sql.SQLite)
	// defer testlib.ResetDatabase()
	app := app.NewApp()

	project, err := app.CreateProject(models.CreateProjectInput{
		Name:        "Testing Project",
		Description: "Testing Description",
	})
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
}

// func TestUpdateLighthouseConfigUpdate(t *testing.T) {
// 	sql.SetupModels(sql.SQLite)
// 	app := NewApp()
//
// 	config, err := app.UpdateLighthouseConfig(models.UpdateLighthouseConfigInput{
// 	})
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, uint(1), config.ID)
// }
//
