package storage

import (
	"io"
)

type Storage interface {
	AddAvatar()
	RemoveAvatar()
	AddOrganizationLogo()
	RemoveOrganizationLogo()
	SaveLighthouseResult(domain string, content io.Reader) error
}
