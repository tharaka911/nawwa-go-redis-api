package initializers

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"context"
	"os"	
)

var ctx = context.Background()

var DB *redis.Client 


func ConnectToRedis(){
	
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	DB = redis.NewClient(opts)

	_, err = DB.Ping(context.Background()).Result()
    if err != nil {
        panic(err)
    }else{
		fmt.Println("Connection to Redis database has been established")
	}
}




