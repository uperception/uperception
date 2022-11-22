package db

import (
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SQLProjectStore struct {
	db *gorm.DB
}

func NewProjectStore() *SQLProjectStore {
	return &SQLProjectStore{
		db: Instance,
	}
}

// Lists all projects
func (s SQLProjectStore) List() ([]*models.Project, error) {
	var projects []*models.Project
	err := s.db.Find(&projects).Error

	return projects, err
}

// Finds the project with the specified ID
func (s SQLProjectStore) FindById(id string) (*models.Project, error) {
	var project models.Project
	err := s.db.Preload(clause.Associations).Preload("LighthouseConfig.Endpoints").Where("id = ?", id).First(&project).Error

	if err != nil {
		return nil, models.ErrNotFound
	}

	return &project, err
}

// Finds the project with the specified ID
func (s SQLProjectStore) FindByName(name string) (*models.Project, error) {
	var project models.Project
	err := s.db.Where("name = ?", name).First(&project).Error

	if err != nil {
		return nil, err
	}

	return &project, err
}

// Saves the project
func (s SQLProjectStore) Save(project *models.Project) error {
	err := s.db.Omit("OrganizationID").Save(project).Error

	return gormErrorInterpreter(err)
}

// Updates the project with the specified ID
func (s SQLProjectStore) Update(project *models.Project) error {
	err := s.db.Save(project).Error

	return err
}

// Deletes the project with the specified ID
func (s SQLProjectStore) Delete(id string) error {
	err := s.db.Delete(&models.Project{}, id).Error

	return err
}
