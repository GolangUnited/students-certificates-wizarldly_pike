package main

import (
	"fmt"
	genpdf2 "gus_certificates/app/certgenerator/genpdf"
	"gus_certificates/app/certgenerator/gentmpl"
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
	res := gentmpl.ParseTemplate(newData)
	fmt.Println(res)
	ok, _ := genpdf2.GeneratePDF(res)
	fmt.Println(ok)
}
