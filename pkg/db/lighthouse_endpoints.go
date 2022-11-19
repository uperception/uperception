package db

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
)

type SQLLighthouseEndpointsStore struct {
	db *gorm.DB
}

func NewLighthouseEndpointsStore() *SQLLighthouseEndpointsStore {
	return &SQLLighthouseEndpointsStore{
		db: Instance,
	}
}

func (s SQLLighthouseEndpointsStore) Save(endpoint *models.LighthouseEndpoint) error {
	err := s.db.Save(endpoint).Error

	return GormErrorInterpreter(err)
}

func (s SQLLighthouseEndpointsStore) Update(endpoint *models.LighthouseEndpoint) error {
	err := s.db.Save(endpoint).Error

	return GormErrorInterpreter(err)
}

func (s SQLLighthouseEndpointsStore) Delete(id string) error {
	return s.db.Delete(models.LighthouseEndpoint{}, id).Error
}
