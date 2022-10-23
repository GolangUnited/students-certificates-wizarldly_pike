package certgenerator

import (
	"fmt"
	"gus_certificates/utils/pdfgenerator"
)

type PdfData struct {
	Template string
	Format   string
	Title    string
	Student  string
	Course   string
	Mentors  string
	Date     string
}

func Starter() error {
	newData := PdfData{
		Template: "sample.html",
		Format:   "A4",
		Title:    "Certificate Golang School",
		Student:  "Khramtsov Denis",
		Course:   "Become a gopher",
		Mentors:  "Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling",
		Date:     "08.09.2022",
	}

	validData, err := validator(newData)
	if err != nil {
		fmt.Println(err.Error())
	}

	buildData, err := build(validData)
	if err != nil {
		fmt.Println(err.Error())
	}

	tmpl, err := ParseTemplate(buildData)
	if err != nil {
		fmt.Println(err.Error())
	}

	buffer, err := pdfgenerator.GeneratePDF(tmpl)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = Save(buffer)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Done!")

	return nil
}
