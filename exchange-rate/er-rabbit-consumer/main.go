package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"er-rabbit-consumer/model"

	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

func main() {
	connDB := fmt.Sprintf("postgres://merlins:root@localhost/er?sslmode=disable")

	db, err := sql.Open("postgres", connDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch bodyFrom(os.Args) {
	case "TRY":
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"TRY", // name
			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
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

		go func() {

			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)

				data := model.Currencies{}
				err := json.Unmarshal(d.Body, &data)
				if err != nil {
					log.Fatal(err)
				}

				insertDynStmt := `insert into currency(base_code, target_code, conversion_rate, created_at) values($1, $2,$3,$4)`
				_, err = db.Exec(insertDynStmt, data.BaseCode, data.TargetCode, data.ConversionRate, data.CreatedAt)
				if err != nil {
					log.Fatal("Database Exec error: ", err)
				}
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	case "USD":
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"USD", // name
			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
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
		go func() {

			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)

				data := model.Currencies{}
				err := json.Unmarshal(d.Body, &data)
				if err != nil {
					log.Fatal(err)
				}

				insertDynStmt := `insert into currency(base_code, target_code, conversion_rate, created_at) values($1, $2,$3,$4)`
				_, err = db.Exec(insertDynStmt, data.BaseCode, data.TargetCode, data.ConversionRate, data.CreatedAt)
				if err != nil {
					log.Fatal("Database Exec error: ", err)
				}
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	case "EUR":
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"EUR", // name
			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
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

		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)

				data := model.Currencies{}
				err := json.Unmarshal(d.Body, &data)
				if err != nil {
					log.Fatal(err)
				}

				insertDynStmt := `insert into currency(base_code, target_code, conversion_rate, created_at) values($1, $2,$3,$4)`
				_, err = db.Exec(insertDynStmt, data.BaseCode, data.TargetCode, data.ConversionRate, data.CreatedAt)
				if err != nil {
					log.Fatal("Database Exec error: ", err)
				}
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	default:
		log.Fatal("Wrong Args")
	}

}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		log.Fatal("give a argument")
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
