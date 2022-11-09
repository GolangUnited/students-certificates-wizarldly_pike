package certgenerator

import (
	"bytes"
	"gopkg.in/validator.v2"
	"html/template"
)

type PdfData struct {
	Title   string `validate:"min=2"`
	Student string `validate:"min=2"`
	Course  string `validate:"min=2"`
	Mentors string `validate:"min=2"`
	Date    string `validate:"min=2"`
}

func New() *PdfData {
	return &PdfData{}
}

func (p *PdfData) SetTitle(s string) *PdfData {
	p.Title = s
	return p
}

func (p *PdfData) SetStudent(s string) *PdfData {
	p.Student = s
	return p
}

func (p *PdfData) SetCourse(s string) *PdfData {
	p.Course = s
	return p
}

func (p *PdfData) SetMentors(s string) *PdfData {
	p.Mentors = s
	return p
}

func (p *PdfData) SetDate(s string) *PdfData {
	p.Date = s
	return p
}

func (p *PdfData) ParseTemplate(input []byte) (*bytes.Buffer, error) {
	t := template.New("certificate")
	tmpl, err := t.Parse((string(input)))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, p)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (p *PdfData) Validate() error {
	err := validator.Validate(p)
	return err
}
