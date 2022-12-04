package app

import (
	"github.com/leometzger/mmonitoring/pkg/models"
)

func (a App) CreateProject(input models.CreateProjectInput) (*models.Project, error) {
	project := models.Project{
		Name:        input.Name,
		Description: input.Description,
	}

	err := a.projectStore.Save(&project)
	if err != nil {
		return nil, err
	}

	return &project, err
}

func (a App) UpdateProject(id string, input models.UpdateProjectInput) (*models.Project, error) {
	project, err := a.projectStore.FindById(id)

	if err != nil {
		return nil, err
	}

	project.Name = input.Name
	project.Description = input.Description
	a.projectStore.Save(project)

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

	return nil
}

func (a App) UpdateLighthouseConfig(id string, input *models.UpdateLighthouseConfigInput) (*models.LighthouseConfig, error) {
	project, err := a.projectStore.FindById(id)
	if err != nil {
		return nil, models.ErrNotFound
	}

	lighthouseConfig := &models.LighthouseConfig{
		ID:        project.LighthouseConfig.ID,
		Enabled:   input.Enabled,
		ProjectID: project.ID,
	}

	err = a.lighthouseConfigStore.Save(lighthouseConfig)
	if err != nil {
		return nil, err
	}

	return lighthouseConfig, nil
}
