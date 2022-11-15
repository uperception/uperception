package app

import (
	"strconv"

	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/rs/zerolog/log"
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
		log.Fatal().Msg("Error consuming the task")
	}

	project, err := a.projectStore.FindById(strconv.FormatUint(uint64(task.ProjectID), 10))
	if err != nil {
		log.Fatal().Msg("Project not found!" + err.Error())
	}

	err = a.lighthouseCollector.Collect(project)
	if err != nil {
		log.Fatal().Msg("Error collecting lighthouse data" + err.Error())
	}
	return err
}

func (a App) SaveLighthouseResult(lighthouseResult *models.LighthouseResult) error {
	return a.lighthouseResultStore.Save(lighthouseResult)
}
