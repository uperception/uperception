package sql

import (
	"fmt"

	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB

type DBType int32

// int mapping
const (
	SQLite   DBType = 0
	Postgres        = 1
)

func SetupModels(dbType DBType) *gorm.DB {
	db, err := getDatabaseByType(dbType)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Organization{})
	db.AutoMigrate(&models.LighthouseConfig{})

	Instance = db

	return Instance
}

func getDatabaseByType(dbType DBType) (*gorm.DB, error) {
	switch dbType {
	case SQLite:
		return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	case Postgres:
		viper.AutomaticEnv()
		pg_user := "mmonitoring"
		pg_password := "12345"
		pg_db := "mmonitoring"
		pg_host := "localhost"
		pg_port := "5432"

		prosgres_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", pg_host, pg_port, pg_user, pg_db, pg_password)

		return gorm.Open(postgres.Open(prosgres_conname), &gorm.Config{})

	default:
		return nil, fmt.Errorf("Database type %d not supported", dbType)
	}
}
