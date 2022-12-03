package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint `gorm:"primary_key"`
	KeycloakID string
	Avatar     string
}

type UserInfo struct {
	Avatar   string
	Email    string
	Name     string
	LastName string
}

type UpdateProfileInput struct {
	Name     string
	LastName string
}
