package app_test

import (
	"testing"

	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/config"
	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/sql"
	"github.com/leometzger/mmonitoring/testlib"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrganization(t *testing.T) {

	sql.SetupModels(sql.SQLite)
	defer testlib.ResetDatabase()
	app := app.NewApp(&config.Config{})

	organization, err := app.CreateOrganization(models.CreateOrganizationInput{
		Name:        "Testing",
		Description: "Testing Description",
	})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), organization.ID)

	// Find
	organization, err = app.FindOrganization("1")
	assert.NoError(t, err)
	assert.Equal(t, organization.Name, "Testing")

	// Update
	organization, err = app.UpdateOrganization("1", models.UpdateOrganizationInput{
		Description: "Another description",
	})

	assert.NoError(t, err)
	assert.Equal(t, organization.Description, "Another description")

	// Delete
	err = app.DeleteOrganization("1")
	assert.NoError(t, err)

	_, err = app.FindOrganization("1")
	assert.Error(t, err)
}
