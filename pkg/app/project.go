package app

import "github.com/leometzger/mmonitoring/pkg/models"

func (a App) CreateProject(input models.CreateProjectInput) (*models.Project, error) {
	project := models.Project{
		Name:        input.Name,
		Description: input.Description,
	}

	a.store.ProjectStore().Save(&project)

	return &project, nil
}

func (a App) UpdateProject(id string, input models.UpdateProjectInput) (*models.Project, error) {
	project, err := a.store.ProjectStore().FindById(id)

	if err != nil {
		return nil, err
	}

	project.Name = input.Name
	project.Description = input.Description
	a.store.ProjectStore().Save(project)

	return project, nil
}

func (a App) QueryProjects() ([]*models.Project, error) {
	projects, err := a.store.ProjectStore().List()

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (a App) FindProject(id string) (*models.Project, error) {
	project, err := a.store.ProjectStore().FindById(id)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (a App) DeleteProject(id string) error {
	err := a.store.ProjectStore().Delete(id)

	if err != nil {
		return err
	}

	return nil
}
