package models

import "gorm.io/gorm"

type SessionState int

const (
	RunningSession SessionState = iota
	FinishedSession
)

type Session struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	State       int
	StartedAt   int64
	FinishedAt  int64
	Path        string
	ProjectID   uint
	SessionPath string
}

type StartSessionInput struct {
	Token string `binding:"required"`
	Path  string `binding:"required"`
}

type StartSessionOutput struct {
	ID uint
}

type PublishEventsInput struct {
	ID         string                   `binding:"required"`
	Path       string                   `binding:"required"`
	Timestamp  uint                     `binding:"required"`
	EventsData []map[string]interface{} `binding:"required"`
}

type FinishSessionInput struct {
	ID uint `binding:"required"`
}
