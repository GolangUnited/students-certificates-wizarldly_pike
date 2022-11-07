package htmltopdf

import (
	"bytes"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type htmltopdf struct {
	pdfgenerator *wkhtmltopdf.PDFGenerator
	bufer        *bytes.Buffer
}

func New() (*htmltopdf, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	htp := &htmltopdf{}
	htp.pdfgenerator = pdfg
	htp.bufer = bytes.NewBuffer(make([]byte, 0))

	pdfg.SetOutput(htp.bufer)

	return htp, nil
}

func (h *htmltopdf) RenderHtmlToPdf(htmlBytes []byte) ([]byte, error) {
	h.bufer.Reset()
	h.pdfgenerator.ResetPages()

	bytesReader := bytes.NewReader(htmlBytes)
	h.pdfgenerator.AddPage(wkhtmltopdf.NewPageReader(bytesReader))
	err := h.pdfgenerator.Create()
	if err != nil {
		return nil, err
	}
	return h.bufer.Bytes(), nil
}
