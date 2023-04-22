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

func (s SQLSessionsStore) FindById(id string) (*models.Session, error) {
	var session models.Session
	err := s.db.Where("id = ?", id).First(&session).Error

	if err != nil {
		return nil, models.ErrNotFound
	}

	return &session, err
}

// Saves the session
func (s SQLSessionsStore) Save(session *models.Session) error {
	err := s.db.Save(session).Error

	return gormErrorInterpreter(err)
}

// Updates the session with the specified ID
func (s SQLSessionsStore) Update(session *models.Session) error {
	err := s.db.Save(session).Error

	return gormErrorInterpreter(err)
}

// Deletes the session with the specified ID
func (s SQLSessionsStore) Delete(id string) error {
	err := s.db.Delete(models.Session{}, id).Error

	return gormErrorInterpreter(err)
}
