package main

import (
	"github.com/tharaka911/go-redis-api/controllers"
	"github.com/tharaka911/go-redis-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToRedis()
	initializers.SyncRedis()
	initializers.InitRabbitMQ()

}

func main() {

	controllers.CatchMessageAndSaveToRedis()

	defer initializers.RabbitMQConn.Close()

}
