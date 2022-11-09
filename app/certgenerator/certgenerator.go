package certgenerator

import (
	"gus_certificates/app/certgenerator/certhtml"
	"gus_certificates/app/types/certdata"
)

type certgenerator interface {
	GenerateCertificate(data *certdata.Data, templateData []byte) ([]byte, error)
}

func NewCertificateHTML() certgenerator {
	return certhtml.New()
}
