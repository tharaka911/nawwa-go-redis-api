package utils

import (
	"github.com/tharaka911/go-redis-api/initializers"
	amqp "github.com/rabbitmq/amqp091-go"

	"log"
	"fmt"
)

var maxChannels = 20
var channelPool = make(chan *amqp.Channel, maxChannels)

func InitChannelPool() {
    for i := 0; i < maxChannels; i++ {
        channel, err := initializers.RabbitMQConn.Channel()
		failOnError(err, "Failed to open a channel")
        channelPool <- channel
    }
	// fmt.Println("length of the channel = ",len(channelPool))
	// fmt.Println(channelPool)
}

func CloseChannelPool() {
	for i := 0; i < maxChannels; i++ {
		channel := <-channelPool
		channel.Close()
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetChannel() (*amqp.Channel, error) {
    // select {
    // case channel := <-channelPool:
    //     return channel, nil
    // default:
    //     // If pool is empty, wait until a channel is returned
    //     channel := <-channelPool
    //     return channel, nil
    // }

	if len(channelPool) > 0 {
		channel := <-channelPool
		fmt.Println("returend channel from the pool")
		return channel, nil
	
		} else {
			fmt.Println("need to wait for a channel to be returned")

		}
		return nil, nil
}

func ReturnChannel(channel *amqp.Channel) {
	channelPool <- channel
}

