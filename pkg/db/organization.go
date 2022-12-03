package db

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
)

type SQLOrganizationStore struct {
	db *gorm.DB
}

func NewOrganizationStore() *SQLOrganizationStore {
	return &SQLOrganizationStore{
		db: Instance,
	}
}

func (s SQLOrganizationStore) List() ([]*models.Organization, error) {
	var organizations []*models.Organization
	err := s.db.Find(&organizations).Preload("LighthouseConfig").Error

	return organizations, gormErrorInterpreter(err)
}

func (s SQLOrganizationStore) FindById(id string) (*models.Organization, error) {
	var organization models.Organization
	err := s.db.Where("id = ?", id).First(&organization).Error

	if err != nil {
		return nil, gormErrorInterpreter(err)
	}

	return &organization, gormErrorInterpreter(err)
}

func (s SQLOrganizationStore) Save(organization *models.Organization) error {
	err := s.db.Save(organization).Error
	return gormErrorInterpreter(err)
}

func (s SQLOrganizationStore) Delete(id string) error {
	err := s.db.Delete(&models.Organization{}, id).Error

	return gormErrorInterpreter(err)
}
