package sqlstore

import (
	"github.com/leometzger/mmonitoring/pkg/db"
	store "github.com/leometzger/mmonitoring/pkg/storage"
	"gorm.io/gorm"
)

type SQLStore struct {
	db *gorm.DB
}

func (s *SQLStore) ProjectStore() store.ProjectStore {
	return SQLProjectStore{
		db: s.db,
	}
}

func (s *SQLStore) SessionStore() store.SessionStore {
	return SQLSessionsStore{
		db: s.db,
	}
}

func (s *SQLStore) OrganizationStore() store.OrganizationStore {
	return SQLOrganizationStore{
		db: s.db,
	}
}

func NewSQLStore() store.Store {
	sqlStore := SQLStore{
		db: db.Instance,
	}

	return &sqlStore
}
