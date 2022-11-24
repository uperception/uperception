package collectors

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"os"
	"os/exec"
	"strconv"

	"github.com/leometzger/mmonitoring/pkg/db"
	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/storage"
)

type LighthouseCollector struct {
	storage storage.Storage
	store   db.LighthouseResultStore
}

func NewLighthouseCollector(storage storage.Storage, store db.LighthouseResultStore) *LighthouseCollector {
	return &LighthouseCollector{
		storage: storage,
		store:   store,
	}
}

// Collects lighthouse data using lighthouse CLI and the configured
// parameters for the endpoint
func (r *LighthouseCollector) Collect(endpoint *models.LighthouseEndpoint) error {
	err := exec.Command(
		"lighthouse",
		endpoint.Url,
		"--chrome-flags='--headless'",
		"--output-path="+getTmpPath(endpoint),
		"--output=json",
	).Run()

	if err != nil {
		return err
	}

	resultFile, err := os.Open(getTmpPath(endpoint))
	if err != nil {
		return err
	}

	file, err := os.ReadFile(getTmpPath(endpoint))
	if err != nil {
		return err
	}

	var result models.LighthouseResult
	err = json.Unmarshal([]byte(file), &result)
	if err != nil {
		return err
	}

	err = r.store.Save(&result)
	if err != nil {
		return err
	}

	err = r.storage.SaveLighthouseResult(endpoint.Url, resultFile)
	if err != nil {
		return err
	}

	return nil
}

// Generates a temporary path to store the JSON result
func getTmpPath(endpoint *models.LighthouseEndpoint) string {
	hasher := sha1.New()
	hasher.Write([]byte(endpoint.Url))
	hasher.Write([]byte(strconv.FormatUint(uint64(endpoint.ID), 10)))

	id := hex.EncodeToString(hasher.Sum(nil))

	return id + ".json"
}
