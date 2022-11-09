package certhtml

import (
	"bytes"
	"gus_certificates/app/types/certdata"
	"html/template"
)

type certhtml struct {
}

func New() *certhtml {
	return &certhtml{}
}

func (c *certhtml) GenerateCertificate(data *certdata.Data, templateData []byte) ([]byte, error) {
	tmpl := template.New("tmpl")

	tmpl, err := tmpl.Parse(string(templateData))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, data.GetDataForTemplate())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
