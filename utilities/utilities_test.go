package utilities

import (
	"os"
	"testing"
)

// Test which mock an http request. Create with POST an element and check with GET
func TestCGetHostAndPort(t *testing.T) {

	expectedHost := "localhost"
	expectedPort := "8080"

	os.Setenv("host", expectedHost)
	os.Setenv("port", expectedPort)

	actualHost, actualPort := GetHostAndPort()

	if expectedHost != actualHost {
		t.Errorf("Test failed expected result: %s actual result: %s", expectedHost, actualHost)
	}
	if expectedPort != actualPort {
		t.Errorf("Test failed expected result: %s actual result: %s", expectedPort, actualPort)
	}

}

// Test which mock an http request. Create with POST an element and check with GET
func TestCGetRabbitInfo(t *testing.T) {

	expectedRabbitURL := "amqp://guest:guest@localhost:5672/"

	os.Setenv("rabbitURL", expectedRabbitURL)

	actualRabbitURL := GetRabbitInfo()

	if expectedRabbitURL != actualRabbitURL {
		t.Errorf("Test failed expected result: %s actual result: %s", expectedRabbitURL, actualRabbitURL)
	}

}
