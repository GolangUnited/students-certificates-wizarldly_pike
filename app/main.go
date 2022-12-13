package main

import (
	"fmt"

	"gus_certificates/app/queu"
)

func main() {
	ch := make(chan queu.CertData, 1)
	go queu.ReceiveData(ch)
	for p := range ch {
		//на основе полученной структуры извлекаем шаблон .html из базы данных заполняем его полями.
		fmt.Println(p)
	}
}
