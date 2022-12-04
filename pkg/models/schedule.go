package models

// CronTab like data structure
type Schedule struct {
	Minute     string
	Hour       string
	DayOfMonth string
	DayWeek    string
	Month      string
}
