// Package controllers contains the main get hanlde function of the project to handle /publish and /consume
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rabbit-exercise/rabbitmq"
)

// HandlePublish Receives and manage publish request 
func HandlePublish(w http.ResponseWriter, r *http.Request) {

	log.Print("HandlePublish: Receiving a /publish request: ")
	exchanges, ok := r.URL.Query()["exchange"]

	if !ok || len(exchanges[0]) < 1 {
		log.Println("Url Param 'exchange' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exchange := exchanges[0]
	messages, ok := r.URL.Query()["message"]

	if !ok || len(messages[0]) < 1 {
		log.Println("Url Param 'message' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	message := messages[0]

	rabbitmqClient := rabbitmq.MakeRabbitClient(rabbitmq.Server)
	err := rabbitmqClient.CreateChannel()
	defer rabbitmqClient.CloseChannel()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = rabbitmqClient.DeclareAndPublishToExchange(exchange, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Reply back to the client request
	w.WriteHeader(http.StatusOK)

}

// HandleConsume Receives and manage consume request 
func HandleConsume(w http.ResponseWriter, r *http.Request) {

	log.Print("HandleConsume: Receiving a /consume request: ")
	queues, ok := r.URL.Query()["queue"]

	if !ok || len(queues[0]) < 1 {
		log.Println("Url Param 'exchange' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queue := queues[0]

	rabbitmqClient := rabbitmq.MakeRabbitClient(rabbitmq.Server)

	err := rabbitmqClient.CreateChannel()
	defer rabbitmqClient.CloseChannel()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	msg, err2 := rabbitmqClient.DeclareAndConsumeFromQueue(queue)

	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	EncodeToJsonWithBody(w, msg)

}

func EncodeToJsonWithBody(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	js, err := json.Marshal(response)
	if err != nil {
		log.Println("Error encoding")
		// returns error
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Write(js)
	return nil
}
