package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:32778")
	failOnError(err, "failed to connecti to mq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to create channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	failOnError(err, "failed to declare exchange")

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	failOnError(err, "failed to declare queue")

	err = ch.QueueBind(q.Name, "", "logs", false, nil)
	failOnError(err, "failed to bind queue to exchange")

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	failOnError(err, "failed to consume from queue")

	for msg := range msgs {
		body := msg.Body
		log.Printf("Received a message %s", body)
		nSec := bytes.Count(body, []byte("."))
		time.Sleep(time.Duration(nSec) * time.Second)
		log.Println("Done")
		msg.Ack(false)
	}
}
