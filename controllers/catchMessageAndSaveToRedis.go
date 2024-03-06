package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/tharaka911/go-redis-api/initializers"
	"github.com/tharaka911/go-redis-api/models"
	"github.com/tharaka911/go-redis-api/utils"
	
)

func CatchMessageAndSaveToRedis() {

	channel, err := utils.GetAChannel()
	utils.FailOnError(err, "Failed to get a channel")

	defer utils.CloseAChannel(channel)

	q, err := channel.QueueDeclare(
		"testnawwa", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {

			go storeInRedis(d.Body)
			// log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func storeInRedis(message []byte) {
	startTime := time.Now()

	var post models.PostRedis

	err := json.Unmarshal(message, &post)
	if err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	postMap := map[string]interface{}{
		"title":         post.Title,
		"body":          post.Body,
		"creation_time": post.CreationTime,
		"updating_time": post.UpdatingTime,
	}

	result := initializers.DB.HMSet(ctx, "posts:"+post.Id, postMap)

	if result.Err() != nil {
		utils.FailOnError(result.Err(), "Failed to store message in Redis")
		return
	} else {
		// Increment the post ID in Redis
		_, err := initializers.DB.Incr(ctx, "post-id").Result()
		if err != nil {
			utils.FailOnError(result.Err(), "Failed to increment post ID in Redis")
			return
		}
	}

	// log.Printf("Received a message with ID %s: %s", post.Id, message)
	// log.Printf("Received a message with ID : %s", post.Id)

	elapsedTime := time.Since(startTime) // Calculate the elapsed time
    // log.Printf("storeInRedis function took %s", elapsedTime)

	log.Printf("Received a message with ID: %s and it stored within %s", post.Id, elapsedTime)
}
