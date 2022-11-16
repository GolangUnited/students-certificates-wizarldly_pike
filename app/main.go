package main

import (
	"fmt"
	c "gus_certificates/app/certgenerator"
	"gus_certificates/utils/pdfgenerator"
	"io/ioutil"
	"log"
)

func main() {
	data := c.New()
	data.SetTitle("Certificate Golang School")
	data.SetStudent("Denis Khramtsov")
	data.SetCourse("Become a gopher")
	data.SetMentors("Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling")
	data.SetDate("08.09.2022")

	err := data.Validate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data validate")

	//Для теста считываем содержимое файла "sample.html":
	template, err := ioutil.ReadFile("sample.html")
	if err != nil {
		log.Fatal(err)
	}

	buffer, err := data.ParseTemplate(template)
	if err != nil {
		log.Fatal(err)
	}

	pdf, err := pdfgenerator.GeneratePDF(buffer)
	if err != nil {
		log.Fatal(err)
	}

	//Для теста сохраняем на диске:
	err = c.Save(pdf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
