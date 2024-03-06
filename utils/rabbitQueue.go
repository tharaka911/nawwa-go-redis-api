package utils

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"context"
	
)

var ctx = context.Background()

func WriteMessageToQueue(message string) {

	// Get a channel from the pool
	channel, err := GetChannel()
	FailOnError(err, "Failed to open a channel")
	// fmt.Println(channel, "channel from the pool")

	defer ReturnChannel(channel)

	// Declare a queue that will be used to send messages
	queue, err := channel.QueueDeclare(
		"testnawwa", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	// Send a message to the queue
	err = channel.PublishWithContext(ctx,
		"",     // exchange
		queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", message)

	//return the channel to the pool
	ReturnChannel(channel)

}


