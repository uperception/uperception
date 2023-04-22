package app

import (
	"fmt"
	"time"

	"github.com/leometzger/mmonitoring/pkg/models"
)

func (a App) StartSession(input models.StartSessionInput) (*models.StartSessionOutput, error) {
	project, err := a.projectStore.FindByToken(input.Token)
	if err != nil {
		return nil, err
	}

	session := models.Session{
		ProjectID: project.ID,
		State:     int(models.RunningSession),
		Path:      input.Path,
		StartedAt: time.Now().Unix(),
	}

	err = a.sessionsStore.Save(&session)
	if err != nil {
		return nil, err
	}

	output := &models.StartSessionOutput{
		ID: session.ID,
	}

	return output, nil
}

func (a App) PublishEvents(input models.PublishEventsInput) error {
	session, err := a.sessionsStore.FindById(input.ID)
	if err != nil {
		return err
	}

	if session.State == int(models.FinishedSession) {
		return models.ErrInvalidEntity
	}

	key := fmt.Sprintf("%s/%d", input.ID, input.Timestamp)
	err = a.storage.AddSessionEvents(key, input.EventsData)
	return err
}

func (a App) FinishSession(input models.FinishSessionInput) error {
	return nil
}
