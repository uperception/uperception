package db

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
)

type SQLSessionsStore struct {
	db *gorm.DB
}

func NewSessionStore() *SQLSessionsStore {
	return &SQLSessionsStore{
		db: Instance,
	}
}

// Saves the session
func (s SQLSessionsStore) Save(session *models.Session) error {
	err := s.db.Save(session).Error

	return err
}

// Updates the session with the specified ID
func (s SQLSessionsStore) Update(session *models.Session) error {
	err := s.db.Save(session).Error

	return err
}

// Deletes the session with the specified ID
func (s SQLSessionsStore) Delete(id string) error {
	err := s.db.Delete(models.Session{}, id).Error

	return err
}
