package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

var animals *[]string

var cmd = &cobra.Command{
	Use:   "consumer",
	Short: "accept greeting from specified kinds of animals",
	Long: `The basic animal description would be: <speed>.<color>.<species>,
but you could use '*'(one any word) or '#'(zero or more any words) as wildcard.
`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := amqp.Dial("amqp://guest:guest@localhost:32778")
		failOnError(err, "failed to dial")

		ch, err := conn.Channel()
		failOnError(err, "failed to create channel")

		err = ch.ExchangeDeclare(
			"animal", // name
			"topic",  // kind
			true,     // durable
			false,    // auto delete
			false,    // internal
			false,    // no wait
			nil,      // args
		)
		failOnError(err, "failed to declare exchange")

		q, err := ch.QueueDeclare(
			"",    //name
			true,  // durable
			false, // autoDelete
			true,  // exclusive
			false, // noWait
			nil)
		failOnError(err, "failed to declare queue")

		for _, animalDesc := range *animals {
			err = ch.QueueBind(
				q.Name,     // name
				animalDesc, // key
				"animal",   // exchange
				false,      // no wait
				nil,        // args
			)
			failOnError(err, "failed to bind queue to exchange")
		}

		consumeCh, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // autoAck
			false,  // exclusive
			false,  // noLocal
			false,  // noWait
			nil,    // args
		)
		failOnError(err, "failed to start to consume")

		for msg := range consumeCh {
			log.Printf("Received: %s", msg.Body)
			msg.Ack(false)
		}
	},
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	animals = cmd.Flags().StringArrayP("animal", "a", []string{}, "animal description: <speed>.<color>.<species>")
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
