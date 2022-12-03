package storage

import (
	"io"
	"net/http"
)

type SignedUrl struct {
	Url    string
	Header http.Header
}

type Storage interface {
	AddAvatar(key string, avatar io.Reader, ext string) (string, error)
	GetAvatarUrl(key string) (*SignedUrl, error)
	RemoveAvatar(key string) error
	SaveLighthouseResult(domain string, content io.Reader) error
}
