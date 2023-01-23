package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func Produce() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Can't connect to rabbitMQ: ", err)
	}
	defer conn.Close()

	//ch, err := conn.Channel()

}
