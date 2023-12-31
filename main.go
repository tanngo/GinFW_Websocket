package main

import "websocket/server"

func main() {
	//server.ExecuteNormalServer() // Call the executeNormalServer() method on the server
	server.ExecuteRabbitMQ() // Call the executeRabbitMQ() method on the server
}
