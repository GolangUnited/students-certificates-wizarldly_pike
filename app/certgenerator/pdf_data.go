package certgenerator

import (
	"bytes"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"html/template"
)

type PdfData struct {
	Title   string
	Student string
	Course  string
	Mentors string
	Date    string
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

func (p PdfData) Validate() error {
	return validation.ValidateStruct(&p,
		// Title cannot be empty
		validation.Field(&p.Title, validation.Required),
		// Student cannot be empty
		validation.Field(&p.Student, validation.Required),
		// Course cannot be empty
		validation.Field(&p.Course, validation.Required),
		// Mentors cannot be empty
		validation.Field(&p.Mentors, validation.Required),
		// Date cannot be empty
		validation.Field(&p.Date, validation.Required),
	)
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
