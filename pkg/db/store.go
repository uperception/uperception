package db

import "github.com/leometzger/mmonitoring/pkg/models"

type ProjectStore interface {
	List() ([]*models.Project, error)
	FindById(string) (*models.Project, error)
	FindByToken(string) (*models.Project, error)
	Save(*models.Project) error
	Update(*models.Project) error
	Delete(string) error
}

type OrganizationStore interface {
	List() ([]*models.Organization, error)
	FindById(string) (*models.Organization, error)
	Save(*models.Organization) error
	Delete(string) error
}

type SessionStore interface {
	Save(*models.Session) error
	FindById(string) (*models.Session, error)
	Update(*models.Session) error
	Delete(string) error
}

type UserStore interface {
	FindByKeycloakId(id string) (*models.User, error)
	Save(*models.User) error
	Update(*models.User) error
	Delete(string) error
}

type LighthouseConfigStore interface {
	Save(*models.LighthouseConfig) error
	Update(*models.LighthouseConfig) error
	Delete(string) error
}

type LighthouseEndpointsStore interface {
	List(uint) ([]*models.LighthouseEndpoint, error)
	FindById(string) (*models.LighthouseEndpoint, error)
	Save(*models.LighthouseEndpoint) error
	SaveBatch([]*models.LighthouseEndpoint) error
	Update(*models.LighthouseEndpoint) error
	Delete(string) error
}

type LighthouseResultStore interface {
	Save(*models.LighthouseResult) error
	Delete(string) error
}

type SessionStorage interface {
	GenerateLink(string) (string, error)
	GetFile(string) ([]byte, error)
}
