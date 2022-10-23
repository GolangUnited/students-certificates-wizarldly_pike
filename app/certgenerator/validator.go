package certgenerator

import (
	"errors"
)

type validData struct {
	Template string
	Format   string
	Title    string
	Student  string
	Course   string
	Mentors  string
	Date     string
	Errors   map[string]string
}

func (p *PdfData) SetTemplate() (string, error) {
	if p.Template == "" {
		return "", errors.New("You need to pass the template data!")
	}
	return p.Template, nil
}

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

func (p *PdfData) SetCourse() (string, error) {
	if p.Course == "" {
		return "", errors.New("You need to pass the course data!")
	}
	return p.Course, nil
}

func (p *PdfData) SetMentors() (string, error) {
	if p.Mentors == "" {
		return "", errors.New("You need to pass the mentors data!")
	}
	return p.Mentors, nil
}

func (p *PdfData) SetDate() (string, error) {
	if p.Format == "" {
		return "", errors.New("You need to pass the date data!")
	}
	return p.Date, nil
}

func validator(pdfData PdfData) (validData, error) {
	vD := validData{
		Errors: make(map[string]string),
	}

	var err error

	vD.Template, err = pdfData.SetTemplate()
	if err != nil {
		vD.Errors["SetTemplate"] = err.Error()
	}

	vD.Format, err = pdfData.SetFormat()
	if err != nil {
		vD.Errors["SetFormat"] = err.Error()
	}

	vD.Title, err = pdfData.SetTitle()
	if err != nil {
		vD.Errors["SetTitle"] = err.Error()
	}

	vD.Student, err = pdfData.SetStudent()
	if err != nil {
		vD.Errors["SetStudent"] = err.Error()
	}

	vD.Course, err = pdfData.SetCourse()
	if err != nil {
		vD.Errors["SetCourse"] = err.Error()
	}

	vD.Mentors, err = pdfData.SetMentors()
	if err != nil {
		vD.Errors["SetMentors"] = err.Error()
	}

	vD.Date, err = pdfData.SetDate()
	if err != nil {
		vD.Errors["SetDate"] = err.Error()
	}

	if len(vD.Errors) == 0 {
		return validData{
			Template: vD.Template,
			Format:   vD.Format,
			Title:    vD.Title,
			Student:  vD.Student,
			Course:   vD.Course,
			Mentors:  vD.Mentors,
			Date:     vD.Date,
			Errors:   vD.Errors,
		}, nil
	} else {
		err = errors.New("Data is not valid!")
		return validData{}, err
	}
}
