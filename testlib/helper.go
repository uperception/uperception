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
}
