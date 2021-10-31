// Package controllers contains the main get hanlde function of the project to handle /service
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rabbit-exercise/rabbitmq"
)

// HandleServiceRequest Receives and manage the request (city and numbers of info)
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

	err := rabbitmq.Client.DeclareAndPublishToExchange(exchange, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Reply back to the client request
	w.WriteHeader(http.StatusOK)

}

// HandleServiceRequest Receives and manage the request (city and numbers of info)
func HandleConsume(w http.ResponseWriter, r *http.Request) {

	log.Print("HandleConsume: Receiving a /consume request: ")
	queues, ok := r.URL.Query()["queue"]

	if !ok || len(queues[0]) < 1 {
		log.Println("Url Param 'exchange' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queue := queues[0]

	msg,err := rabbitmq.Client.DeclareAndConsumeFromQueue(queue)

	if err != nil {
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
