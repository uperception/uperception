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

func (s SQLLighthouseEndpointsStore) List(lighthouseConfigID uint) ([]*models.LighthouseEndpoint, error) {
	var endpoints []*models.LighthouseEndpoint
	err := s.db.Find(&endpoints).Where("LighthouseConfigID = ?", lighthouseConfigID).Error

	return endpoints, gormErrorInterpreter(err)
}

func (s SQLLighthouseEndpointsStore) FindById(id string) (*models.LighthouseEndpoint, error) {
	var endpoint *models.LighthouseEndpoint
	err := s.db.Where("id = ?", id).First(&endpoint).Error

	return endpoint, gormErrorInterpreter(err)
}

func (s SQLLighthouseEndpointsStore) Save(endpoint *models.LighthouseEndpoint) error {
	err := s.db.Save(endpoint).Error

	return gormErrorInterpreter(err)
}

func (s SQLLighthouseEndpointsStore) SaveBatch(endpoints []*models.LighthouseEndpoint) error {
	err := s.db.CreateInBatches(endpoints, 100).Error

	return gormErrorInterpreter(err)
}

func (s SQLLighthouseEndpointsStore) Update(endpoint *models.LighthouseEndpoint) error {
	err := s.db.Save(endpoint).Error

	return gormErrorInterpreter(err)
}

func (s SQLLighthouseEndpointsStore) Delete(id string) error {
	err := s.db.Delete(models.LighthouseEndpoint{}, id).Error

	return gormErrorInterpreter(err)
}
