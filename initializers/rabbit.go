package initializers

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"log"
	"os"
)

var RabbitMQConn *amqp.Connection

func InitRabbitMQ() error {

	rabbit_uri := os.Getenv("RABBIT_URL")
	var err error
	RabbitMQConn, err = amqp.Dial(rabbit_uri)
	failOnError(err, "Failed to connect to RabbitMQ")

	if RabbitMQConn == nil {
		log.Panic("Failed to connect to RabbitMQ")
	} else {
		log.Println("RabbitMQ connection established")
	}

	return err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
