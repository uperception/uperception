package sql

import (
	"gorm.io/gorm"
)

type SQLStore struct {
	db *gorm.DB
}

func (s *SQLStore) ProjectStore() ProjectStore {
	return SQLProjectStore{
		db: s.db,
	}
}

func (s *SQLStore) SessionStore() SessionStore {
	return SQLSessionsStore{
		db: s.db,
	}
}

func (s *SQLStore) OrganizationStore() OrganizationStore {
	return SQLOrganizationStore{
		db: s.db,
	}
}

func NewSQLStore() Store {
	sqlStore := SQLStore{
		db: Instance,
	}

	return &sqlStore
}
