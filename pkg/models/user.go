package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Name     string
	LastName string
	Email    string
}
