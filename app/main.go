package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gus_certificates/utils/storage/s3"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found!")
	}
}

func main() {
	endpoint, _ := os.LookupEnv("ENDPOINT")
	accessKey, _ := os.LookupEnv("ACCESS_KEY")
	secretKey, _ := os.LookupEnv("SECRET_KEY")

	storage, err := s3.NewStorage(endpoint, accessKey, secretKey)
	if err != nil {
		log.Fatal(err)
	}

	buffer, err := storage.GetTemplate("sample.html")
	if err != nil {
		log.Fatal(err)
	}

	buffer, err = storage.GetCertificate("example.pdf")
	if err != nil {
		log.Fatal(err)
	}

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

	body, err = ioutil.ReadFile("example.pdf")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	err = storage.SaveCertificate("example2.pdf", body)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.DeleteCertificate("example2.pdf")
	if err != nil {
		log.Fatal(err)
	}

	err = storage.DeleteTemplate("sample2.html")
	if err != nil {
		log.Fatal(err)
	}
}
