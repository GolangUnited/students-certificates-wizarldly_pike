package queu

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type CertData struct {
	CourseName       string `json:"course_name"`
	CourseType       string `json:"course_type"`
	CourseHours      string `json:"course_hours"`
	CourseDate       string `json:"course_date"`
	CourseMentors    string `json:"course_mentors"`
	StudentFirstname string `json:"student_firstname"`
	StudentLastname  string `json:"student_lastname"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ReceiveData(data chan CertData) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	bch := make(chan CertData, 1)
	go func() {
		for d := range msgs {
			var p CertData
			err := json.Unmarshal(d.Body, &p)
			if err != nil {
				fmt.Errorf("Unmarshal data error!")
			}
			bch <- p
		}
	}()
	for k := range bch {
		data <- k
	}

	<-forever
}
