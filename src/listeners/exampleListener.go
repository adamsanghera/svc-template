package listeners

import (
	"github.com/streadway/amqp"
	"log"
	"fmt"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func handleMessage(msg []byte) {
	log.Printf("Received Message %s", msg)
}

func ExampleListener(conn *amqp.Connection) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel for ExampleListener")
	defer ch.Close()

	msgs, err := ch.Consume(
		"hello",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")

	go func() {
		for d := range msgs {
			handleMessage(d.Body)
		}
	}()
}