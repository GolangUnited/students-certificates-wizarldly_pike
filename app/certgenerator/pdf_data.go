package certgenerator

import (
	"bytes"
	"errors"
)

type PdfData struct {
	//Template string
	Format  string
	Title   string
	Student string
	course  string
	mentors string
	Date    string
}

func New() *PdfData {
	return &PdfData{}
}

func (p *PdfData) SetCourse(s string) *PdfData {
	p.course = s
	return p
}

func (p *PdfData) SetMentors(s string) *PdfData {
	p.mentors = s
	return p
}

//func (p *PdfData) SetTemplate() (string, error) {
//	if p.Template == "" {
//		return "", errors.New("You need to pass the template data!")
//	}
//	return p.Template, nil
//}

func (p *PdfData) SetFormat() (string, error) {
	if p.Format == "" {
		return "", errors.New("You need to pass the format data!")
	}
	return p.Format, nil
}

func (p *PdfData) SetTitle() (string, error) {
	if p.Title == "" {
		return "", errors.New("You need to pass the title data!")
	}
	return p.Title, nil
}

func (p *PdfData) SetStudent() (string, error) {
	if p.Student == "" {
		return "", errors.New("You need to pass the student data!")
	}
	return p.Student, nil
}

//func (p *PdfData) SetMentors() (string, error) {
//	if p.mentors == "" {
//		return "", errors.New("You need to pass the mentors data!")
//	}
//	return p.mentors, nil
//}

func (p *PdfData) SetDate() (string, error) {
	if p.Format == "" {
		return "", errors.New("you need to pass the date data")
	}
	return p.Date, nil
}

func (p *PdfData) ParseTemplate(templateBytes bytes.Buffer) (bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, p)
	return buf, nil
}

func (p *PdfData) Validate() error {
	return nil
}

//func Starter() error {
//	newData := PdfData{
//		Template: "sample.html",
//		Format:   "A4",
//		Title:    "Certificate Golang School",
//		Student:  "Khramtsov Denis",
//		course:   "Become a gopher",
//		mentors:  "Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling",
//		Date:     "08.09.2022",
//	}
//
//	validData, err := validator(newData)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	buildData, err := build(validData)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	tmpl, err := ParseTemplate(buildData)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	buffer, err := pdfgenerator.GeneratePDF(tmpl)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	err = Save(buffer)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	fmt.Println("Done!")
//
//	return nil
//}
