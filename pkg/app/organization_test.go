package app

import (
	"testing"

	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrganization(t *testing.T) {
	db.SetupModels(db.SQLite)
	// defer testlib.ResetDatabase()
	app := NewApp()

	organization, err := app.CreateOrganization(models.CreateOrganizationInput{
		Name:        "Testing",
		Description: "Testing Description",
	})

	assert.NoError(t, err)
	assert.Equal(t, uint(1), organization.ID)
}
