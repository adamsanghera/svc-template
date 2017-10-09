package main

import (
	bus "./src"
	"fmt"
)

func handleExample() {
	fmt.println("Omg we just caught an example message!")
}

func main() {
	// Config for connections.
	exampleQueue := bus.QueueDeclaration{
		"example-queue", // name
		false, 			// durable
		false, 			// auto-delete
		false, 			// exclusive
		false, 			// noWait
		nil, 			// args
	}
	bus.AddQueue(exampleQueue)

	bus.Init()
	// Make connections.
	// Listen!
}