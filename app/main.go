package main

import (
	"gus_certificates/app/certgenerator"
	"log"
)

func main() {
	err := certgenerator.Starter()
	if err != nil {
		log.Fatal(err)
	}
}
