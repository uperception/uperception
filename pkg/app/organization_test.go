package app

import (
	"testing"

	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/sql"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrganization(t *testing.T) {
	sql.SetupModels(sql.SQLite)
	// defer testlib.ResetDatabase()
	app := NewApp()

	organization, err := app.CreateOrganization(models.CreateOrganizationInput{
		Name:        "Testing",
		Description: "Testing Description",
	})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), organization.ID)
}
