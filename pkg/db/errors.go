package db

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/leometzger/mmonitoring/pkg/models"
	"gorm.io/gorm"
)

// Gorm oriented error interpreter.
// It interpret gorm errors and pgerrors
func GormErrorInterpreter(err error) error {
	if err == nil {
		return err
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return models.ErrDuplicated
		}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ErrNotFound
	}

	if errors.Is(err, gorm.ErrInvalidData) || errors.Is(err, gorm.ErrInvalidField) || errors.Is(err, gorm.ErrInvalidValue) || errors.Is(err, gorm.ErrInvalidValueOfLength) {
		return models.ErrInvalidEntity
	}

	return err
}
