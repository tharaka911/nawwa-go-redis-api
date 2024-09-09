package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tharaka911/go-redis-api/initializers"
	"github.com/tharaka911/go-redis-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToRedis()
	initializers.SyncRedis()
}

func main() {

	r := routes.SetupRouter()
	
	var port = os.Getenv("PORT")

	fmt.Println("Server started on port " + port)
	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, r))

	fmt.Println("Server started on port " + port)
//checking for conflicts


}
