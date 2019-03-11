package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var levels *[]string
var cmd = &cobra.Command{
	Use:   "sub",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:32778")
		failOnError(err, "failed to connecti to mq")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "failed to create channel")
		defer ch.Close()

		err = ch.ExchangeDeclare("logs", "direct", true, false, false, false, nil)
		failOnError(err, "failed to declare exchange")

		q, err := ch.QueueDeclare("", false, false, true, false, nil)
		failOnError(err, "failed to declare queue")

		for _, bindkey := range *levels {
			err = ch.QueueBind(q.Name, bindkey, "logs", false, nil)
			failOnError(err, "failed to bind queue to exchange")
		}

		msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
		failOnError(err, "failed to consume from queue")

		for msg := range msgs {
			body := msg.Body
			log.Printf("Received a message %s", body)
			msg.Ack(false)
		}
	},
}

func init() {
	levels = cmd.Flags().StringSliceP("level", "l", []string{"debug"}, "log level to accept, one of [debug, info, warn, error]")
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
