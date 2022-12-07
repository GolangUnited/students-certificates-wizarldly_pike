package htmltopdf

import (
	"bytes"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type htmltopdf struct {
	pdfgenerator *wkhtmltopdf.PDFGenerator
}

func New() (*htmltopdf, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	htp := &htmltopdf{}
	htp.pdfgenerator = pdfg

	return htp, nil
}

func (h *htmltopdf) RenderHtmlToPdf(htmlBytes []byte) ([]byte, error) {
	h.pdfgenerator.ResetPages()
	h.pdfgenerator.Buffer().Reset()

	bytesReader := bytes.NewReader(htmlBytes)
	pageReader := wkhtmltopdf.NewPageReader(bytesReader)
	pageReader.Encoding.Set("utf-8")
	h.pdfgenerator.AddPage(pageReader)
	err := h.pdfgenerator.Create()
	if err != nil {
		return nil, err
	}
	return h.pdfgenerator.Bytes(), nil
}
