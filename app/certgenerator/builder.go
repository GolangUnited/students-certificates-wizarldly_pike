package certgenerator

import (
	"errors"
	"fmt"
)

func build(vD validData) (PdfData, error) {
	var buildData PdfData
	if len(vD.Errors) == 0 {
		buildData.Template = vD.Template
		buildData.Format = vD.Format
		buildData.Title = vD.Title
		buildData.Student = vD.Student
		buildData.Course = vD.Course
		buildData.Mentors = vD.Mentors
		buildData.Date = vD.Date
		return buildData, nil
	} else {
		err := errors.New("Building error!")
		for key, value := range vD.Errors {
			fmt.Printf("%s : %s \n", key, value)
		}
		return buildData, err
	}
}
