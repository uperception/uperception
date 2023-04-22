package app

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/scheduler"
)

func (a App) CreateProject(input models.CreateProjectInput) (*models.Project, error) {
	project := models.Project{
		Name:        input.Name,
		Description: input.Description,
		Token:       "UP-" + strconv.Itoa(1000000+rand.Intn(9000000)),
		LighthouseConfig: models.LighthouseConfig{
			Enabled:   true,
			Endpoints: []models.LighthouseEndpoint{},
			Schedule: models.LighthouseSchedule{
				Minute:     "*",
				Hour:       "12",
				DayOfMonth: "*",
				DayWeek:    "?",
				Month:      "*",
			},
		},
	}

	err := a.projectStore.Save(&project)
	if err != nil {
		return nil, err
	}

	a.createProjectSchedule(&project)

	return &project, err
}

func (a App) UpdateProject(id string, input models.UpdateProjectInput) (*models.Project, error) {
	project, err := a.projectStore.FindById(id)

	if err != nil {
		return nil, err
	}

	project.Name = input.Name
	project.Description = input.Description
	project.LighthouseConfig.Schedule = input.LighthouseConfig.Schedule
	project.LighthouseConfig.Enabled = input.LighthouseConfig.Enabled
	err = a.projectStore.Save(project)

	if err != nil {
		return nil, err
	}

	a.createProjectSchedule(project)

	return project, nil
}

func (a App) QueryProjects() ([]*models.Project, error) {
	projects, err := a.projectStore.List()

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (a App) FindProject(id string) (*models.Project, error) {
	project, err := a.projectStore.FindById(id)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (a App) DeleteProject(id string) error {
	err := a.projectStore.Delete(id)
	if err != nil {
		return err
	}

	a.scheduler.DeleteSchedule(fmt.Sprintf("lighthouse-%v", id))
	return nil
}

func (a App) FindProjectByToken(token string) (*models.Project, error) {
	project, err := a.projectStore.FindByToken(token)

	return project, err
}

func (a App) createProjectSchedule(project *models.Project) error {
	schedule := models.Schedule{
		Minute:     project.LighthouseConfig.Schedule.Minute,
		Hour:       project.LighthouseConfig.Schedule.Hour,
		DayOfMonth: project.LighthouseConfig.Schedule.DayOfMonth,
		DayWeek:    project.LighthouseConfig.Schedule.DayWeek,
		Month:      project.LighthouseConfig.Schedule.Month,
	}
	message := models.RunLighthouseProjectMessage{ProjectID: project.ID}
	jsonMsg, _ := json.Marshal(message)

	err := a.scheduler.Schedule(schedule, scheduler.ScheduleConfig{
		Id:      fmt.Sprintf("lighthouse-%v", project.ID),
		Payload: string(jsonMsg),
	})

	return err
}
