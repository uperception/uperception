package models

import (
	"gorm.io/gorm"
)

type Environment uint8
type LighthouseState uint8

const (
	Mobile Environment = iota
	Desktop
	Tablet
)

const (
	Created LighthouseState = iota
	Queued
	Running
	None
)

// Lighthouse specific schedule model
type LighthouseSchedule struct {
	ID                 uint `json:"-"`
	LighthouseConfigID uint
	Minute             string `binding:"required"`
	Hour               string `binding:"required"`
	DayOfMonth         string `binding:"required"`
	DayWeek            string `binding:"required"`
	Month              string `binding:"required"`
}

// Lighthouse configuration
type LighthouseConfig struct {
	ID        uint `gorm:"primary_key" json:"-"`
	ProjectID uint `json:"-"`
	Enabled   bool
	Schedule  LighthouseSchedule
	Endpoints []LighthouseEndpoint `json:"-"`
}

type LighthouseEndpoint struct {
	ID                 uint `gorm:"primary_key"`
	Enabled            bool
	LighthouseConfigID uint   `gorm:"index:,unique,composite:urlpathing" json:"-"`
	Url                string `gorm:"index:,unique,composite:urlpathing"`
	Header             string
	LighthouseState    LighthouseState
}

type LighthouseResult struct {
	gorm.Model
	LighthouseConfigID uint             `json:"-"`
	ID                 uint             `gorm:"primary_key"`
	GatherMode         string           `json:"gatherMode"`
	Environment        Environment      `json:"-"`
	Audits             LighthouseAudits `json:"audits"`
}

type LighthouseAudits struct {
	ID                     uint             `gorm:"primary_key" json:"-"`
	LighthouseResultID     uint             `gorm:"index:idx_result"`
	FirstContentfulPaint   LighthouseMetric `json:"first-contentful-paint"`
	FirstMeaningfulPaint   LighthouseMetric `json:"first-meaningful-paint"`
	LargestContentfulPaint LighthouseMetric `json:"largest-contentful-paint"`
	SpeedIndex             LighthouseMetric `json:"speed-index"`
	TotalBlockingTime      LighthouseMetric `json:"total-blocking-time"`
	MaxPotentialFid        LighthouseMetric `json:"max-potential-fid"`
	TimeToInteractive      LighthouseMetric `json:"time-to-interactive"`
	NetworkRoundTripTimes  LighthouseMetric `json:"network-round-trip-times"`
	NetworkServerLatency   LighthouseMetric `json:"network-server-latency"`
}

type LighthouseMetric struct {
	ID                 uint    `gorm:"primary_key" json:"-"`
	LighthouseAuditsID uint    `gorm:"index:idx_audits" json:"-"`
	Title              string  `json:"title"`
	Unit               string  `json:"unit"`
	Score              float32 `json:"score"`
	Value              float32 `json:"value"`
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
	Enabled  bool               `binding:"required"`
	Schedule LighthouseSchedule `binding:"required"`
}

type LighthouseEndpointInput struct {
	ID      uint
	Enabled bool
	Url     string `binding:"required"`
	Header  string
}
