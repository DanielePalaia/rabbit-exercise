// This package create a manage the rabbitmq connection to the rabbitmq server
package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitServer struct {
	connString string
	conn       *amqp.Connection
}

var (
	Server *RabbitServer
)

func MakeRabbitServer(connString string) *RabbitServer {
	server := new(RabbitServer)
	server.connString = connString
	return server
}

func (server *RabbitServer) Connect() {

	conn, err := amqp.Dial(server.connString)
	server.failOnErrorFatal(err, "Failed to connect to RabbitMQ")
	server.conn = conn

}

func (server *RabbitServer) failOnErrorFatal(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
