package storage

import (
	"gus_certificates/utils/storage/local"
)

type storage interface {
	GetTemplate(string) ([]byte, error)
	GetCertificate(string) ([]byte, error)
	GetCertificateLink(string) (string, error)

	SaveTemplate(string, []byte) error
	SaveCertificate(string, []byte) error

	DeleteTemplate(string) error
	DeleteCertificate(string) error
}

func NewLocal() (storage, error) {
	return local.New()
}
