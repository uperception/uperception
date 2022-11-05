package queue

import (
	"github.com/leometzger/mmonitoring-runner/models"
)

type Task struct {
	Id      string
	Project *models.Project
}

type Queue interface {
	GetTask() (*Task, error)
	FinishTask(id string) error
}
