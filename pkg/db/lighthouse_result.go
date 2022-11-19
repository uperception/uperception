package db

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
)

type SQLLighthouseResultStore struct {
	db *gorm.DB
}

func NewLighthouseResultStore() *SQLLighthouseResultStore {
	return &SQLLighthouseResultStore{
		db: Instance,
	}
}

func (s SQLLighthouseResultStore) Save(result *models.LighthouseResult) error {
	return s.db.Save(result).Error
}

func (s SQLLighthouseResultStore) Delete(id string) error {
	return s.db.Delete(models.LighthouseResult{}, id).Error
}
