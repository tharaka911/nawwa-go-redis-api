package main

import (
	"github.com/tharaka911/go-redis-api/initializers"
	"github.com/tharaka911/go-redis-api/routes"
	"log"
	"net/http"
	"os"
	
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToRedis()
	initializers.SyncRedis()
}

func main() {

	r := routes.SetupRouter()
	
	var port = os.Getenv("PORT")
	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, r))

}
