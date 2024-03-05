package main

import (
	"github.com/tharaka911/go-redis-api/initializers"
	"github.com/tharaka911/go-redis-api/routes"
	"github.com/tharaka911/go-redis-api/utils"
	"log"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToRedis()
	initializers.SyncRedis()
	initializers.InitRabbitMQ()
	utils.InitChannelPool()

	
}

func main() {

	r := routes.SetupRouter()

	var port = os.Getenv("PORT")
	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, r))

	// Send a message to the queue
	


	//close the rabbitmq connection
	defer utils.CloseChannelPool()
	defer initializers.RabbitMQConn.Close()
}
