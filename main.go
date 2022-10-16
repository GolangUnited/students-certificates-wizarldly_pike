package main

import (
	"gus_certificates/app/certgenerator/genpdf"
	"gus_certificates/app/certgenerator/gentmpl"
	"gus_certificates/app/certgenerator/savepdf"
	"log"
)

func main() {
	newData := gentmpl.PdfData{
		Templ:   "sample.html",
		Format:  "A4",
		Title:   "Certificate Golang School",
		Student: "Khramtsov Denis",
		Course:  "Become a gopher",
		Mentors: "Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling",
		Date:    "08.09.2022",
	}

	res, err := gentmpl.ParseTemplate(newData)
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := genpdf.GeneratePDF(res)
	if err != nil {
		log.Fatal(err)
	}

	err = savepdf.Save(pdfg)
	if err != nil {
		log.Fatal(err)
	}
}
