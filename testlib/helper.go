package testlib

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/sql"
)

func ResetDatabase() {
	sql.Instance.Delete(&models.Organization{})
	sql.Instance.Delete(&models.User{})
	sql.Instance.Delete(&models.Session{})
	sql.Instance.Delete(&models.Project{})

	sql.Instance.Delete(&models.LighthouseConfig{})
	sql.Instance.Delete(&models.LighthouseResult{})
	sql.Instance.Delete(&models.LighthouseMetric{})
	sql.Instance.Delete(&models.LighthouseEndpoint{})
}
