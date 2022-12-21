package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gus_certificates/app/certgenerator"
	"gus_certificates/utils/pdfgenerator/htmltopdf"
	"gus_certificates/utils/storage/local"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gus_certificates/app/queue"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found!")
	}
}

func main() {

	amqpHost, ok := os.LookupEnv("AMQP_HOST")
	if !ok {
		log.Println("No amqpHost data!")
	}

	ctx := context.Background()
	d := time.Now().Add(60 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, d)

	defer cancel()

	data := make(chan []byte)

	go queue.ReceiveData(ctx, data, amqpHost)

	loop := true

	var p certgenerator.CertData

	for loop {
		select {
		case k := <-data:

			err := json.Unmarshal(k, &p)
			if err != nil {
				log.Fatal(err)
			}

			certGen := certgenerator.New(p)

			err = certGen.ValidateData()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("ok!")

			storage, err := local.New()
			if err != nil {
				log.Fatal(err)
			}

			tmpl, err := storage.GetTemplate("template")
			if err != nil {
				log.Fatal(err)
			}

			certHTML, err := certGen.GenerateCertHTML(tmpl)
			if err != nil {
				log.Fatal(err)
			}

			pdfgen, err := htmltopdf.New()
			if err != nil {
				log.Fatal(err)
			}

			pdf, err := pdfgen.RenderHtmlToPdf(certHTML)
			if err != nil {
				log.Fatal(err)
			}

			err = storage.SaveCertificate("file name", pdf)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Finish!")

		case <-ctx.Done():
			loop = false
		}
	}
}
