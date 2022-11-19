package models

import "gorm.io/gorm"

type LighthouseState uint8

const (
	Created LighthouseState = iota
	Running
)

type Project struct {
	gorm.Model
	ID               uint   `gorm:"primary_key"`
	Name             string `gorm:"index:idx_name,unique"`
	Description      string
	OrganizationID   uint
	UserID           uint
	LighthouseConfig LighthouseConfig
	LighthouseState  LighthouseState
	Sessions         []Session
}

type CreateProjectInput struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
}

type UpdateProjectInput struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
}
