package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint `gorm:"primary_key"`
	KeycloakID string
}

type UserInfo struct {
	Email    string
	Name     string
	LastName string
	// Add props
}

type UpdateProfileInput struct {
	Name     string
	LastName string
}
