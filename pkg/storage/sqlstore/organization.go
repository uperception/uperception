package sqlstore

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
)

type SQLOrganizationStore struct {
	db *gorm.DB
}

func (s SQLOrganizationStore) List() ([]*models.Organization, error) {
	var organizations []*models.Organization
	err := s.db.Find(&organizations).Error

	return organizations, err
}

func (s SQLOrganizationStore) FindById(id string) (*models.Organization, error) {
	var organization models.Organization
	err := s.db.Where("id = ?", id).First(&organization).Error

	if err != nil {
		return nil, err
	}

	return &organization, err
}

func (s SQLOrganizationStore) Save(organization *models.Organization) error {
	err := s.db.Save(organization).Error
	return err
}

func (s SQLOrganizationStore) Delete(id string) error {
	err := s.db.Delete(models.Organization{}, id).Error

	return err
}
