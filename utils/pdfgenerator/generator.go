package pdfgenerator

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePDF(buf *bytes.Buffer) (*bytes.Buffer, error) {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(buf))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Buffer(), nil
}
