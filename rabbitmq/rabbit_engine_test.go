package rabbitmq

import (
	"testing"

	"rabbit-exercise/utilities"
)

// Test HandlePublish is working fine
func TestHDeclareAndPublishToExchange(t *testing.T) {

	exchange := "exchangeTestHDeclareAndPublishToExchange"
	queue := "queueTestHDeclareAndPublishToExchange"
	expectedMessage := "message1"

	client := initialize()

	client.DeclareAndPublishToExchange(exchange, expectedMessage)
	client.DeclareQueue(queue)
	client.BindQueue(exchange, queue)
	msg, _ := client.ConsumeFromQueue(queue)

	if msg != expectedMessage {
		t.Errorf("wrong message received: got %v want %v",
			msg, expectedMessage)
	}

}

// Setting up rabbitmq cluster, exchange and queue
func initialize() *rabbitClient {

	rabbitURL := utilities.GetRabbitInfo()
	client := MakeRabbitClient(rabbitURL)
	client.Connect()

	return client

}
