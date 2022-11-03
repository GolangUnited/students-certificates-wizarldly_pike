package certgenerator

import (
	"errors"
)

type validData struct {
	Template string `json:"template", validate:"minLen=3"`
	Format   string
	Title    string
	Student  string
	Course   string
	Mentors  string
	Date     string
	Errors   map[string]error
}

func validator(pdfData PdfData) (validData, error) {
	vD := validData{
		Errors: make(map[string]string),
	}

	var err error

	vD.Template, err = pdfData.SetTemplate()
	if err != nil {
		vD.Errors["SetTemplate"] = err
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
