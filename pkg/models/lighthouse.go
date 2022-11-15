package models

import (
	"gorm.io/gorm"
)

type Environment uint8

const (
	Mobile Environment = iota
	Desktop
	Tablet
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
	ID                 uint `gorm:"primary_key"`
	LighthouseConfigID uint
	Environment        Environment
	GatherMode         string           `json:"gatherMode"`
	Audits             LighthouseAudits `json:"audits"`
}

type LighthouseAudits struct {
	ID                     uint             `gorm:"primary_key"`
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
	gorm.Model
	ID                 uint    `gorm:"primary_key"`
	LighthouseAuditsID uint    `gorm:"index:idx_audits"`
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
	Enabled     bool                      `binding:"required"`
	Periodicity uint8                     `binding:"required"`
	Endpoints   []LighthouseEndpointInput `binding:"required"`
}

type LighthouseEndpointInput struct {
	ID     uint
	Url    string `binding:"required"`
	Header string
}
