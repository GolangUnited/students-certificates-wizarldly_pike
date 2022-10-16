package savepdf

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func Save(pdfgen *wkhtmltopdf.PDFGenerator) error {
	err := pdfgen.WriteFile("./utils/storage/local/certificates/" + "certificate.pdf")
	if err != nil {
		return err
	}
	return nil
}
