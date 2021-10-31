// Package utilities provide general utilities functions which can be reused in other projects
package utilities

import (
	"os"
)

// GetHostAndPort returns the host and the port to bind the service defined in the OS variables host and port
func GetHostAndPort() (string, string) {

	host := os.Getenv("host")
	port := os.Getenv("port")

	return host, port

}

//Get rabbitmq informations: connection string and the exchangename
func GetRabbitInfo() string {
	rabbitURL := os.Getenv("rabbitURL")

	return rabbitURL

}
