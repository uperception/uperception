package models

import "gorm.io/gorm"

// User session representation
type Session struct {
	gorm.Model
	ID         uint `gorm:"primary_key"`
	State      string
	StartedAt  int64
	FinishedAt int64
	Path       string
	ProjectID  uint
}
