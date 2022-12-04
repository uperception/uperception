package scheduler

import "github.com/leometzger/mmonitoring/pkg/models"

type Scheduler interface {
	Schedule(id string, schedule models.Schedule) error
	EnableSchedule(id string) error
	DisableSchedule(id string) error
}
