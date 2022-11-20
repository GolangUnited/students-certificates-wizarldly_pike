package storage

import (
	"gus_certificates/utils/storage/local"
)

type Storage interface {
	GetTemplate(string) ([]byte, error)
	GetCertificate(string) ([]byte, error)
	GetCertificatePath(string) (string, error)

	SaveTemplate(string, []byte) error
	SaveCertificate(string, []byte) error

	DeleteTemplate(string) error
	DeleteCertificate(string) error
}

func NewLocal() (Storage, error) {
	return local.New()
}
