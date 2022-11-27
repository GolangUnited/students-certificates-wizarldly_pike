package main

import (
	"fmt"
	"gus_certificates/utils/storage/s3"
	"io/ioutil"
	"log"
)

func main() {
	endpoint := "localhost:9000"
	accessKey := "..." //необходимо вписать для проверки
	secretKey := "..." //необходимо вписать для проверки
	storage, err := s3.NewStorage(endpoint, accessKey, secretKey)
	if err != nil {
		log.Fatal(err)
	}
	buffer, err := storage.GetTemplate("sample.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer)

	buffer, err = storage.GetCertificate("example.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer)

	url, err := storage.GetCertificatePath("example.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
	body, err := ioutil.ReadFile("sample.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	err = storage.SaveTemplate("sample2.html", body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok1!")

	body, err = ioutil.ReadFile("example.pdf")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	err = storage.SaveCertificate("example2.pdf", body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok2!")
	err = storage.DeleteCertificate("example2.pdf")
	if err != nil {
		log.Fatal(err)
	}
	err = storage.DeleteTemplate("sample2.html")
	if err != nil {
		log.Fatal(err)
	}
}
