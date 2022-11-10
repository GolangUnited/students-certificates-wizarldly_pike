package pdfgenerator

import "gus_certificates/utils/pdfgenerator/htmltopdf"

type PdfGenerator interface {
	RenderHtmlToPdf([]byte) ([]byte, error)
}

func New() (PdfGenerator, error) {
	return htmltopdf.New()
}
