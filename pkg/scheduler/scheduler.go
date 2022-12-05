package scheduler

import "github.com/leometzger/mmonitoring/pkg/models"

type ScheduleConfig struct {
	Id      string
	Payload string
}

type Scheduler interface {
	Schedule(schedule models.Schedule, config ScheduleConfig) error
	DeleteSchedule(id string) error
	EnableSchedule(id string) error
	DisableSchedule(id string) error
}
