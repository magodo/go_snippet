package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var cmd = &cobra.Command{
	Use:   "producer animal1 [animal2 [...]] ",
	Short: "publish kinds of animal, in format like: <speed>.<color>.<species>",
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

		for _, key := range args {
			body := fmt.Sprintf("hello i'm %s", key)
			err = ch.Publish(
				"animal",
				key,
				false, // mandatory
				false, // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			failOnError(err, "failed to publish")
			log.Printf("Sent: %s", body)
		}
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
