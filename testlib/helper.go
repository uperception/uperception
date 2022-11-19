package testlib

import (
	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/models"
)

func ResetDatabase() {
	db.Instance.Delete(&models.Organization{})
	db.Instance.Delete(&models.User{})
	db.Instance.Delete(&models.Session{})
	db.Instance.Delete(&models.Project{})
	db.Instance.Delete(&models.LighthouseConfig{})
	db.Instance.Delete(&models.LighthouseResult{})
	db.Instance.Delete(&models.LighthouseMetric{})
	db.Instance.Delete(&models.LighthouseEndpoint{})
}

func CreateApp() *app.App {
	return nil
}
