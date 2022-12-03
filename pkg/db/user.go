package db

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
)

type SQLUserStore struct {
	db *gorm.DB
}

func NewUserStore() *SQLUserStore {
	return &SQLUserStore{
		db: Instance,
	}
}

// Finds the project with the specified ID
func (s SQLUserStore) FindByKeycloakId(keycloakId string) (*models.User, error) {
	var user models.User
	err := s.db.Where("keycloak_id = ?", keycloakId).First(&user).Error
	if err != nil {
		return nil, gormErrorInterpreter(err)
	}

	return &user, gormErrorInterpreter(err)
}

// Saves the project
func (s SQLUserStore) Save(user *models.User) error {
	err := s.db.Save(user).Error

	return gormErrorInterpreter(err)
}

// Updates the project with the specified ID
func (s SQLUserStore) Update(user *models.User) error {
	err := s.db.Save(user).Error

	return gormErrorInterpreter(err)
}

// Deletes the project with the specified ID
func (s SQLUserStore) Delete(id string) error {
	err := s.db.Delete(&models.User{}, id).Error

	return gormErrorInterpreter(err)
}
