package collectors

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/sql"
	"github.com/leometzger/mmonitoring/pkg/storage"
	"github.com/rs/zerolog/log"
)

type LighthouseCollector struct {
	storage storage.Storage
	store   sql.LighthouseResultStore
}

func NewLighthouseCollector(storage storage.Storage, store sql.LighthouseResultStore) *LighthouseCollector {
	return &LighthouseCollector{
		storage: storage,
		store:   store,
	}
}

// Collects lighthouse data
func (r *LighthouseCollector) Collect(project *models.Project) error {
	for i, endpoint := range project.LighthouseConfig.Endpoints {
		err := exec.Command(
			"lighthouse",
			endpoint.Url,
			"--chrome-flags='--headless'",
			"--output-path="+getTmpPath(i),
			"--output=json",
		).Run()

		if err != nil {
			log.Fatal().Msg("Error running lighthouse command for" + endpoint.Url + err.Error())
		}
	}

	for i, endpoint := range project.LighthouseConfig.Endpoints {
		resultFile, err := os.Open(getTmpPath(i))
		if err != nil {
			log.Err(err).Msg("Error opening lighthouse result file")
			return err
		}

		file, err := ioutil.ReadFile(getTmpPath(i))
		if err != nil {
			log.Err(err).Msg("Error reading lighthouse file")
			return err
		}

		var result models.LighthouseResult
		err = json.Unmarshal([]byte(file), &result)
		if err != nil {
			log.Err(err).Msg("Error unmarshaling lighthouse result")
			return err
		}

		err = r.store.Save(&result)
		if err != nil {
			log.Err(err).Msg("Error saving lighthouse result into database")
		}

		err = r.storage.SaveLighthouseResult(endpoint.Url, resultFile)
		if err != nil {
			log.Err(err).Msg("Error saving lighthouse result into storage")
			return err
		}
	}

	return nil
}

func getTmpPath(index int) string {
	return "mmonitoring-" + strconv.FormatInt(int64(index), 10) + ".json"
}
