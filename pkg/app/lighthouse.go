package app

import (
	"strconv"

	"github.com/leometzger/mmonitoring/pkg/models"
)

func (a App) EnqueueLighthouseTask(project *models.Project) error {
	return a.queue.Publish(project.ID)
}

func (a App) EnqueueAllLighthouseTasks() error {
	return nil
}

func (a App) CollectLighthouseData() error {
	task, err := a.queue.GetTask()
	if err != nil {
		return err
	}

	if task == nil {
		return nil
	}

	project, err := a.projectStore.FindById(strconv.FormatUint(uint64(task.ProjectID), 10))
	if err != nil {
		return err
	}

	err = a.lighthouseCollector.Collect(project)
	if err != nil {
		return err
	}

	err = a.queue.FinishTask(task.Id)
	if err != nil {
		return err
	}

	return err
}
func (a App) SaveLighthouseResult(lighthouseResult *models.LighthouseResult) error {
	return a.lighthouseResultStore.Save(lighthouseResult)
}

func (a App) ListLighthouseEndpoints(projectId string) ([]*models.LighthouseEndpoint, error) {
	project, err := a.projectStore.FindById(projectId)
	if err != nil {
		return nil, err
	}

	endpoints, err := a.lighthouseEndpointsStore.List(project.LighthouseConfig.ID)
	return endpoints, err
}

func (a App) FindLighthouseEndpoint(id string) (*models.LighthouseEndpoint, error) {
	endpoint, err := a.lighthouseEndpointsStore.FindById(id)
	return endpoint, err
}

func (a App) CreateLighthouseEndpoint(projectId string, input models.LighthouseEndpointInput) (*models.LighthouseEndpoint, error) {
	project, err := a.projectStore.FindById(projectId)
	if err != nil {
		return nil, err
	}

	endpont := &models.LighthouseEndpoint{
		ID:                 input.ID,
		LighthouseConfigID: project.LighthouseConfig.ID,
		Url:                input.Url,
		Header:             input.Header,
		LighthouseState:    models.Created,
	}
	err = a.lighthouseEndpointsStore.Save(endpont)
	if err != nil {
		return nil, err
	}

	return endpont, nil
}

func (a App) CreateLighthouseEndpoints(endpoint []models.LighthouseEndpointInput) ([]*models.LighthouseEndpoint, error) {
	return nil, nil
}

func (a App) UpdateLighthouseEndpoint(id string, input models.LighthouseEndpointInput) (*models.LighthouseEndpoint, error) {
	endpoint, err := a.lighthouseEndpointsStore.FindById(id)
	if err != nil {
		return nil, err
	}

	endpoint.Header = input.Header
	endpoint.Url = input.Url
	err = a.lighthouseEndpointsStore.Save(endpoint)
	return endpoint, err
}

func (a App) DeleteLighthouseEndpoint(id string) error {
	return a.lighthouseEndpointsStore.Delete(id)
}
