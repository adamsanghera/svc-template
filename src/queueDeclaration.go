package src

import "github.com/streadway/amqp"

type QueueDeclaration struct {
	Name 		string
	Durable 	bool
	AutoDelete 	bool
	Exclusive 	bool
	NoWait 		bool
	Args 		amqp.Table
}