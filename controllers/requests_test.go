package controllers

import (
	"net/http"
	"net/http/httptest"
	"rabbit-exercise/rabbitmq"
	"rabbit-exercise/utilities"
	"testing"
)

// Test HandlePublish is working fine
func TestHandlePublishOK(t *testing.T) {

	initialize()

	// Get the element
	message := "?exchange=TestHandlePublishOK&message=hello"
	reqGet := MockGetRequest(t, message)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePublish)
	handler.ServeHTTP(rr, reqGet)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestHandlePublishFailure(t *testing.T) {

	initialize()
	// Get the element
	message := "?wrong=hello"
	reqGet := MockGetRequest(t, message)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleConsume)
	handler.ServeHTTP(rr, reqGet)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// Check that we are consuming the message we publish using HandleConsume
func TestHandleConsumeOK(t *testing.T) {

	expectedMessage := "\"hello\""
	initialize()
	// Get the element
	message := "?queue=TestHandleConsumeOK"
	reqGet := MockGetRequest(t, message)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleConsume)
	handler.ServeHTTP(rr, reqGet)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if rr.Body.String() != expectedMessage {
		t.Errorf("wrong message received: got %v want %v",
			rr.Body.String(), expectedMessage)
	}

}

// Setting up rabbitmq cluster, exchange and queue
func initialize() {

	rabbitURL := utilities.GetRabbitInfo()
	rabbitmq.Server = rabbitmq.MakeRabbitServer(rabbitURL)
	rabbitmq.Server.Connect()
	client := rabbitmq.MakeRabbitClient(rabbitmq.Server)
	client.CreateChannel()
	defer client.CloseChannel()
	client.DeclareAndPublishToExchange("TestHandlePublishOK", "hello")
	client.DeclareQueue("TestHandleConsumeOK")
	client.BindQueue("TestHandlePublishOK", "TestHandleConsumeOK")
}

// Mock a get request
func MockGetRequest(t *testing.T, message string) *http.Request {
	req, err := http.NewRequest("GET", "/rabbit/"+message, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}
