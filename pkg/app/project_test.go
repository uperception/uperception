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
	projectID := strconv.FormatUint(uint64(project.ID), 10)

	assert.NoError(t, err)

	// Find
	project, err = app.FindProject(projectID)
	assert.NoError(t, err)
	assert.Equal(t, project.Name, "Testing Project")

	// Update
	project, err = app.UpdateProject(projectID, models.UpdateProjectInput{
		Name:        "Testing Project v2",
		Description: "Testing Description v2",
	})

	assert.NoError(t, err)
	assert.Equal(t, project.Name, "Testing Project v2")

	// Delete
	err = app.DeleteProject(projectID)
	assert.NoError(t, err)

	_, err = app.FindProject(projectID)
	assert.Error(t, err)
}
