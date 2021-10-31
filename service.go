package main

import (
	"log"
	"net/http"
	"rabbit-exercise/rabbitmq"
	"rabbit-exercise/utilities"
)

// Main entry point of the service listening on 8080
func main() {

	host, port := utilities.GetHostAndPort()
	rabbitURL := utilities.GetRabbitInfo()
	rabbitmq.Client = rabbitmq.MakeRabbitClient(rabbitURL)
	rabbitmq.Client.Connect()
	createServer(host, port)

}

func createServer(host string, port string) {
	// routes defined in routes.go
	router := NewRouter()
	if err := http.ListenAndServe(host+":"+port, router); err != nil {
		log.Fatalf("critical error listing and serving on host %v and port %v: error %v", host, port, err)
	}

}
