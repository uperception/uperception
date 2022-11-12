package collectors

import (
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/storage"
)

type LighthouseCollector struct {
	storage storage.Storage
}

func NewLighthouseCollector(storage storage.Storage) *LighthouseCollector {
	return &LighthouseCollector{
		storage: storage,
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
			log.Println("Error running lighthouse command for", endpoint.Url, err)
			break
		}
	}

	for i, endpoint := range project.LighthouseConfig.Endpoints {
		resultFile, err := os.Open(getTmpPath(i))
		if err != nil {
			log.Fatal(err)
		}

		err = r.storage.SaveLighthouseResult(endpoint.Url, resultFile)
		if err != nil {
			log.Fatal("Error saving lighthouse result", err)
		}
	}

	return nil
}

func getTmpPath(index int) string {
	return "mmonitoring-" + strconv.FormatInt(int64(index), 10) + ".json"
}
