package app

import (
	"github.com/leometzger/mmonitoring/pkg/models"
)

func (a App) CollectLighthouseData() error {
	task, err := a.queue.GetTask()
	if err != nil {
		return err
	}

	if task == nil {
		return nil
	}

	endpointID := task.Body
	endpoint, err := a.lighthouseEndpointsStore.FindById(endpointID)
	if err != nil {
		return err
	}

	err = a.lighthouseCollector.Collect(endpoint)
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

	// body := strconv.FormatUint(uint64(endpont.ID), 10)
	// _, err = a.queue.Publish(&body)
	// if err != nil {
	// 	fmt.Println(err.Error()) // TODO: Log this better
	// }

	return endpont, nil
}

func (a App) CreateLighthouseEndpoints(projectID string, input []models.LighthouseEndpointInput) ([]*models.LighthouseEndpoint, error) {
	var endpoints []*models.LighthouseEndpoint

	project, err := a.projectStore.FindById(projectID)
	if err != nil {
		return nil, err
	}

	for _, in := range input {
		endpoints = append(endpoints, &models.LighthouseEndpoint{
			Url:                in.Url,
			Header:             in.Header,
			LighthouseConfigID: project.LighthouseConfig.ID,
			LighthouseState:    models.Created,
		})
	}

	err = a.lighthouseEndpointsStore.SaveBatch(endpoints)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
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
