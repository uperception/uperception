package db

import "github.com/leometzger/mmonitoring/pkg/models"

type ProjectStore interface {
	List() ([]*models.Project, error)
	FindById(id string) (*models.Project, error)
	Save(project *models.Project) error
	Update(project *models.Project) error
	Delete(id string) error
}

type OrganizationStore interface {
	List() ([]*models.Organization, error)
	FindById(id string) (*models.Organization, error)
	Save(organization *models.Organization) error
	Delete(id string) error
}

type SessionStore interface {
	Save(project *models.Session) error
	Update(project *models.Session) error
	Delete(id string) error
}

type UserStore interface {
	Save(project *models.User) error
	Update(project *models.User) error
	Delete(id string) error
}

type LighthouseConfigStore interface {
	Save(config *models.LighthouseConfig) error
	Update(config *models.LighthouseConfig) error
	Delete(id string) error
}

type LighthouseResultStore interface {
	Save(result *models.LighthouseResult) error
	Delete(id string) error
}

type SessionStorage interface {
	GenerateLink(key string) (string, error)
	GetFile(key string) ([]byte, error)
}
