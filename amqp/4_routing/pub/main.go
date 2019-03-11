package main

import (
	"log"
	"os"
	"strings"

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

	err = ch.ExchangeDeclare("logs", "direct", true, false, false, false, nil)
	failOnError(err, "failed to declare exchange")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs",
		severityFrom(os.Args),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "failed to publish message")
	log.Printf("[x] Sent %s", body)
}

func parseLog(log string) (level, msg string) {
	splits := strings.Split(log, "|")
	if len(splits) != 2 {
		level = "debug"
		msg = log
	} else {
		level = splits[0]
		msg = splits[1]
	}
	return
}

func bodyFrom(args []string) string {
	if len(args) == 1 {
		return "empty message"
	}
	_, msg := parseLog(strings.Join(args[1:], " "))
	return msg
}

func severityFrom(args []string) string {
	if len(args) == 1 {
		return "debug"
	}
	level, _ := parseLog(strings.Join(args[1:], " "))
	return level
}
