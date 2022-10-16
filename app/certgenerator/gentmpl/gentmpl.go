package gentmpl

import (
	"bytes"
	"html/template"
)

type PdfData struct {
	Templ   string
	Format  string
	Title   string
	Student string
	Course  string
	Mentors string
	Date    string
}

type RequestPdf struct {
	Buf    *bytes.Buffer
	Format string
}

func SetPdf(buf *bytes.Buffer, format string) RequestPdf {
	return RequestPdf{
		Buf:    buf,
		Format: format,
	}
}

func ParseTemplate(p PdfData) (RequestPdf, error) {
	tmpl, err := template.ParseFiles("./utils/storage/local/templates/" + p.Templ)
	if err != nil {
		return RequestPdf{}, err
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, p)
	if err != nil {
		return RequestPdf{}, err
	}
	return SetPdf(buf, p.Format), nil
}
