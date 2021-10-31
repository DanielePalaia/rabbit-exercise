package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type rabbitClient struct {
	connString string
	ch         *amqp.Channel
}

var (
	Client *rabbitClient
)

func MakeRabbitClient(connString string) *rabbitClient {
	client := new(rabbitClient)
	client.connString = connString

	return client
}

func (client *rabbitClient) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (client *rabbitClient) Connect() {

	conn, err := amqp.Dial(client.connString)
	client.failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	client.failOnError(err, "Failed to open a channel")
	client.ch = ch
	//defer ch.Close()

}

func (client *rabbitClient) DeclareAndPublishToExchange(exchange string, message string) {

	err := client.ch.ExchangeDeclare(
		exchange, // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	client.failOnError(err, "Failed to declare an exchange")

	err = client.ch.Publish(
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	client.failOnError(err, "Failed to publish to the exchange")

}

func (client *rabbitClient) DeclareQueue(queue string) {
	// Declare a queue
	_, err := client.ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	client.failOnError(err, "Failed to declare a queue")
}

func (client *rabbitClient) ConsumeFromQueue(queue string) string {

	msg, _, err := client.ch.Get(queue, true)
	output := string(msg.Body)

	client.failOnError(err, "Failed to declare a queue")

	return output
}

func (client *rabbitClient) DeclareAndConsumeFromQueue(queue string) string {

	// Declare a queue
	client.DeclareQueue(queue)

	output := client.ConsumeFromQueue(queue)

	return output

}

// Binds a queue to an exchange
func (client *rabbitClient) BindQueue(exchange string, queue string) {
	err := client.ch.QueueBind(
		queue,    // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	client.failOnError(err, "Failed to bind the queue to the exchange")

}
