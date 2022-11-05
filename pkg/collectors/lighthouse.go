package collectors

import (
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/leometzger/mmonitoring-runner/storage"
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
func (r *LighthouseCollector) Collect(urls []string) error {
	for i, url := range urls {
		err := exec.Command(
			"lighthouse",
			url,
			"--chrome-flags='--headless'",
			"--output-path="+getTmpPath(i),
			"--output=json",
		).Run()

		if err != nil {
			log.Println("Error running lighthouse command for", url, err)
			break
		}
	}

	for i, url := range urls {
		resultFile, err := os.Open(getTmpPath(i))
		if err != nil {
			log.Fatal(err)
		}

		err = r.storage.SaveLighthouseResult(url, resultFile)
		if err != nil {
			log.Fatal("Error saving lighthouse result", err)
		}
	}

	return nil
}

func getTmpPath(index int) string {
	return "mmonitoring-" + strconv.FormatInt(int64(index), 10) + ".json"
}
