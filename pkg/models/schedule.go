package models

import (
	"strings"
)

// CronTab like data structure
type Schedule struct {
	Minute     string
	Hour       string
	DayOfMonth string
	DayWeek    string
	Month      string
}

// Transform a Schedule to a cron expression
// year should always be * in the uperception context
// we do not have to worry about yearly schedules
func (s Schedule) ToCronExpression() string {
	cronParts := []string{s.Minute, s.Hour, s.DayOfMonth, s.Month, s.DayWeek, "*"}

	for i, part := range cronParts {
		if part == "" {
			cronParts[i] = "*"
		}
	}

	return strings.Join(cronParts, " ")
}

func (s Schedule) IsValid() bool {
	return true
}
