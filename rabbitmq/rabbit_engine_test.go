// test for rabbitmq operations
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

	server := initialize()

	client := MakeRabbitClient(server)
	client.CreateChannel()
	defer client.CloseChannel()

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
func initialize() *RabbitServer {

	rabbitURL := utilities.GetRabbitInfo()
	server := MakeRabbitServer(rabbitURL)
	server.Connect()

	return server

}
