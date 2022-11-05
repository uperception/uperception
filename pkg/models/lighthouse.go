package models

import "gorm.io/gorm"

type LighthouseConfig struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Enabled     bool
	Periodicity uint8
	Endpoints   []LighthouseEndpoint
	ProjectID   uint
}

type LighthouseEndpoint struct {
	gorm.Model
	ID                 uint `gorm:"primary_key"`
	Url                string
	Header             string
	LighthouseConfigID uint
}

type UpdateLighthouseConfigInput struct {
	Enabled     bool                 `binding:"required"`
	Periodicity uint8                `binding:"required"`
	Endpoints   []LighthouseEndpoint `binding:"required"`
}
