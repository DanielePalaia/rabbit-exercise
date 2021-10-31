# rabbitmq-exercise

Simple exercise for the rabbitmq platform team interview

## Requirements

Design and code a simple API to do basic operations on a
RabbitMQ server, using any web server library of your choice
- The web server API must meet the following criteria: </br>
- The API must expose an endpoint to publish a message to the exchange. </br>
This endpoint must handle the creation of any necessary objects to publish a
message.
- The API must expose an endpoint to consume a message from the queue. </br>


## Design
The service has been written in Golang using gorilla as router for the entrypoints. </br>
It supports two entrypoints: Get operations on /publish and /consume where: </br></br>
- **/publish**: takes in input two parameters, the exchange to publish and the message to publish and the request will look like this </br>
http://localhost:8088/publish?exchange=testrabbitmq&message=test </br>
It returns http ok in case everything goes fine otherwise an http error message </br></br>
- **/consume**: takes in input the queue to consume. Just one message will be read from the queue if exists and the request will look like this: </br>
http://localhost:8088/consume?queue=rabbitclient </br>
It returns the message read encoded as json </br>

### Input parameter
The software takes the input parameters as OS environment variables in order to be ready to be deployed on Docker/K8s if necessary.
These are the env variables to set:
- **export host=localhost**
- **export port=8088**
- **export rabbitURL=amqp://daniele:daniele@localhost:5672/**

### Packages
The software is written taken modularity in mind: There are a set of packages defined
- **main:** The main package. Here the service get started and entrypoints are set with gorilla
- **controllers:** Here the handler functions are defined to manager requests and responses
- **rabbitmq:** Here are defined all functions related to rabbit: declares, create of queues ecc
- **utilities:** Some utilities functions

### How to build
Modules have been used. to build should be necessary to  </br>
**go build -o service**  </br>
on the root folder  </br>

### Unit testing
Unit testing have been provided in order to provide coverage of the functions: net/http/httptest has been used in order to mock and test the entrypoints
To run tests is necessary to have a rabbitmq started locally (ex using docker)

```
dpalaia@dpalaia-a02 vmw-rabbit-exercise % go test ./...
?       vmw-vdp [no test files]
ok      vmw-vdp/controllers     0.740s
?       vmw-vdp/rabbitmq        [no test files]
ok      vmw-vdp/utilities       (cached)
```
### Deployment
Binary can be called after compilation and after having set input OS variables.</br>
They can be found in ./bin directory for osx and linux </br>
A Dockerfile has been provided also in order to deploy it in Docker and eventually on K8s

The Dockerfile can be used to create an image like: </br>
**docker build -t rabbitmq-exercise .** </br>
And then you can use docker run -e to specifiy env variables and run the image.</br>
Maybe also put the image on Dockerhub and deploy the project in k8s</br>

