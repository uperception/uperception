package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID               uint   `gorm:"primary_key"`
	Name             string `gorm:"index:idx_name,unique"`
	Description      string
	OrganizationID   uint
	UserID           uint
	Token            string
	LighthouseConfig LighthouseConfig
	Sessions         []Session
}

type CreateProjectInput struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
}

type UpdateProjectInput struct {
	Name             string `binding:"required"`
	Description      string `binding:"required"`
	LighthouseConfig LighthouseConfig
}
