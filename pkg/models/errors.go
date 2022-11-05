package models

import "errors"

var ErrDuplicated = errors.New("duplicated error")
var ErrInvalidEntity = errors.New("invalid entity")
var ErrCannotBeDeleted = errors.New("cannot be deleted")
