package storage

type Storage interface {
	SaveLighthouseResult(domain string) error
	StoreMetadata() error
}
