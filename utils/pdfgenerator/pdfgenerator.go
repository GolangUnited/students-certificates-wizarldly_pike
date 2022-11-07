package pdfgenerator

import "gus_certificates/utils/pdfgenerator/htmltopdf"

type pdfgenerator interface {
	RenderHtmlToPdf([]byte) ([]byte, error)
}

func New() (pdfgenerator, error) {
	return htmltopdf.New()
}
