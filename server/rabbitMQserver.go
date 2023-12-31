package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

var upgraderRBMQ = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ExecuteRabbitMQ() {
	//conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"testing",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		ws, err := upgraderRBMQ.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer ws.Close()
		for msg := range msgs {
			ws.WriteMessage(websocket.TextMessage, msg.Body)
		}
	})
	router.Run(":7007")
}
