package src

import (
	"github.com/streadway/amqp"

	"fmt"
	"log"
)

var queues []QueueDeclaration = []QueueDeclaration{}
var connection amqp.Connection

func AddQueue(queue QueueDeclaration) {
	queues = append(queues, queue)
}

func AddListener(handleFunction fn, queueName string) {
	channel, err := connection.Channel()
	channel.Consume()
}

func Init() {
	// Connect to the MQ.
	connection = connectToQueue("amqp://guest:guest@localhost:5672/")
	defer connection.Close()

	// Open a channel for the purpose of declaring queues.
	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel for queue declaration.")
	defer channel.Close()

	// Declare all of the queues... Maybe we should do "lazy" declares?
	// Where we don't declare until we listen?
	qArr := declareQueues(channel, queues)

	// Set up subscriptions to the ones we like.
	fmt.Println(qArr)
}

// Helper function.  Used to check the return value of each call.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// Wraps QueueDeclare, to take our struct as an argument.
func declareQueue(dec QueueDeclaration, channel *amqp.Channel) (amqp.Queue, error) {
	q, err := channel.QueueDeclare(
		dec.Name,
		dec.Durable,
		dec.AutoDelete,
		dec.Exclusive,
		dec.NoWait,
		dec.Args)
	return q, err
}

func connectToQueue(queueURL string) *amqp.Connection {
	connection, err := amqp.Dial(queueURL)
	failOnError(err, "Failed to connect to RabbitMQ.")
	return connection
}

func declareQueues(channel *amqp.Channel, queues[]QueueDeclaration) []amqp.Queue {
	qArr := []amqp.Queue{}
	for i := 0; i < len(queues) ; i += 1 {
		q, err := declareQueue(queues[i], channel)
		failOnError(err, "Failed to declare queue " + queues[i].Name)
		qArr = append(qArr, q)
	}
	return qArr
}