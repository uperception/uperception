package api

import (
	"net/http"

	"github.com/leometzger/mmonitoring/pkg/models"
)

func getStatus(err error) int {
	switch err {
	case models.ErrCannotBeDeleted:
	case models.ErrDuplicated:
		return http.StatusConflict

	case models.ErrNotFound:
		return http.StatusNotFound

	case models.ErrInvalidEntity:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
