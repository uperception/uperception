package storage

import (
	"io"
)

type Storage interface {
	AddAvatar(key string, avatar io.Reader) error
	GetAvatarUrl(key string) (string, error)
	RemoveAvatar(key string) error
	SaveLighthouseResult(domain string, content io.Reader) error
}
