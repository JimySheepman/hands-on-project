package task

import (
	"bytes"
	"encoding/json"
	"er-api-consumer/model"
	"log"

	"github.com/streadway/amqp"
)

func Send(data model.Currencies, name string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        buf.Bytes(),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %v\n", data)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
