package db

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
)

type SQLLighthouseConfigStore struct {
	db *gorm.DB
}

func NewLighthouseConfigStore() *SQLLighthouseConfigStore {
	return &SQLLighthouseConfigStore{
		db: Instance,
	}
}

func (s SQLLighthouseConfigStore) Save(config *models.LighthouseConfig) error {
	return s.db.Save(config).Error
}

func (s SQLLighthouseConfigStore) Update(config *models.LighthouseConfig) error {
	err := s.db.Save(config).Error

	return err
}

func (s SQLLighthouseConfigStore) Delete(id string) error {
	return s.db.Delete(models.LighthouseConfig{}, id).Error
}
