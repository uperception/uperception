package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Logo        string
	Name        string
	Description string
	Projects    []Project
	Users       []User `gorm:"many2many:organization_users;"`
}

type CreateOrganizationInput struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Logo        string
}

type UpdateOrganizationInput struct {
	Description string `binding:"required"`
	Logo        string
}

type AddUserToOrgInput struct {
	OrganizationId uint `binding:"required"`
	UserId         uint `binding:"required"`
}
