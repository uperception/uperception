package collectors

import (
	"log"
	"os"
	"os/exec"

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
func (r *LighthouseCollector) Collect(url string) error {
	log.Println("Start running for", url)
	err := exec.Command(
		"lighthouse",
		url,
		"--chrome-flags='--headless'",
		"--output-path=/tmp/reports/mmonitoring.json",
		"--output=json",
	).Run()
	log.Println("End running")

	if err != nil {
		log.Fatal("Error running lighthouse command", err)
	}

	log.Println("Saving /temp monitoring")
	resultFile, err := os.Open("/tmp/reports/mmonitoring.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Adding to storage")
	err = r.storage.SaveLighthouseResult(url, resultFile)
	if err != nil {
		log.Fatal("Error saving lighthouse result", err)
	}
	log.Println("Adding to storage")

	return nil
}
