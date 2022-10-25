package storage

import "io"

type Storage interface {
	SaveLighthouseResult(domain string, content io.Reader) error
}
