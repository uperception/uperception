package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID             uint   `gorm:"primary_key"`
	Name           string `gorm:"index:idx_name,unique"`
	Token          string
	Description    string
	OrganizationID uint
	Sessions       []Session
}

type LighthouseProject struct {
	Domain string
	Urls   []string
}

type CreateProjectInput struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
}

type UpdateProjectInput struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
}
