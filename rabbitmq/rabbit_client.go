// This package manages channels and rabbitmq client opration
package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type rabbitClient struct {
	connString string
	server     *RabbitServer
	ch         *amqp.Channel
}

func MakeRabbitClient(server *RabbitServer) *rabbitClient {
	client := new(rabbitClient)
	client.server = server

	return client
}

func (client *rabbitClient) failOnErrorNonFatal(err error, msg string) error {
	if err != nil {
		log.Printf("%s: %s", msg, err)
		client.ch.Close()
	}
	return err
}

func (client *rabbitClient) CreateChannel() error {

	ch, err := client.server.conn.Channel()
	client.failOnErrorNonFatal(err, "Failed to open a channel")
	client.ch = ch
	return err

}

func (client *rabbitClient) DeclareAndPublishToExchange(exchange string, message string) error {

	err := client.ch.ExchangeDeclare(
		exchange, // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	client.failOnErrorNonFatal(err, "Failed to declare an exchange")
	if err != nil {
		return err
	}

	err = client.ch.Publish(
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	client.failOnErrorNonFatal(err, "Failed to publish to the exchange")
	return err

}

func (client *rabbitClient) DeclareQueue(queue string) error {
	// Declare a queue
	_, err := client.ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	client.failOnErrorNonFatal(err, "Failed to declare a queue")
	return err
}

func (client *rabbitClient) ConsumeFromQueue(queue string) (string, error) {

	msg, _, err := client.ch.Get(queue, true)
	output := string(msg.Body)

	client.failOnErrorNonFatal(err, "Failed to get items from the queue")

	return output, err
}

func (client *rabbitClient) DeclareAndConsumeFromQueue(queue string) (string, error) {

	// Declare a queue
	client.DeclareQueue(queue)

	output, err := client.ConsumeFromQueue(queue)

	return output, err

}

// Binds a queue to an exchange
func (client *rabbitClient) BindQueue(exchange string, queue string) error {
	err := client.ch.QueueBind(
		queue,    // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	client.failOnErrorNonFatal(err, "Failed to bind the queue to the exchange")
	return err

}

// Binds a queue to an exchange
func (client *rabbitClient) CloseChannel() error {
	err := client.ch.Close()
	client.failOnErrorNonFatal(err, "Failed to close a channel")
	return err

}
