package testlib

import (
	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/models"
)

func ResetDatabase() {
	db.Instance.Delete(&models.Organization{})
	db.Instance.Delete(&models.User{})
	db.Instance.Delete(&models.Session{})
	db.Instance.Delete(&models.Project{})
}
