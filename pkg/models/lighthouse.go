package models

import (
	"gorm.io/gorm"
)

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

type LighthouseResult struct {
	gorm.Model
	ID                     uint `gorm:"primary_key"`
	LighthouseConfigID     uint
	FirstContentfulPaint   LighthouseMetric
	FirstMeaningfulPaint   LighthouseMetric
	LargestContentfulPaint LighthouseMetric
	SpeedIndex             LighthouseMetric
	TotalBlockingTime      LighthouseMetric
	MaxPotentialFid        LighthouseMetric
	TimeToInteractive      LighthouseMetric
	NetworkRoundTripTimes  LighthouseMetric
	NetworkServerLatency   LighthouseMetric
}

type LighthouseMetric struct {
	gorm.Model
	ID                 uint `gorm:"primary_key"`
	LighthouseResultID uint
	Title              string
	Unit               string
	Score              float32
	Value              float32
}

type CreateLighthouseResultInput struct {
	FirstContentfulPaint   LighthouseMetric `binding:"required"`
	FirstMeaningfulPaint   LighthouseMetric `binding:"required"`
	LargestContentfulPaint LighthouseMetric `binding:"required"`
	SpeedIndex             LighthouseMetric `binding:"required"`
	TotalBlockingTime      LighthouseMetric `binding:"required"`
	MaxPotentialFid        LighthouseMetric `binding:"required"`
	TimeToInteractive      LighthouseMetric `binding:"required"`
	NetworkRoundTripTimes  LighthouseMetric `binding:"required"`
	NetworkServerLatency   LighthouseMetric `binding:"required"`
}

type UpdateLighthouseConfigInput struct {
	Enabled     bool                      `binding:"required"`
	Periodicity uint8                     `binding:"required"`
	Endpoints   []LighthouseEndpointInput `binding:"required"`
}

type LighthouseEndpointInput struct {
	Url    string `binding:"required"`
	Header string
}
