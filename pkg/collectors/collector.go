package collectors

import "github.com/leometzger/mmonitoring/pkg/models"

type Collector interface {
	Collect(endpoint *models.LighthouseEndpoint) error
}

type Crawler interface {
	CrawlUrls(domain string, root string) ([]string, error)
}
