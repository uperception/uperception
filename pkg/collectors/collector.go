package collectors

import "github.com/leometzger/mmonitoring/pkg/models"

type Collector interface {
	Collect(project models.Project) error
}
