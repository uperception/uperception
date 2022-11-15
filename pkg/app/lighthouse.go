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
