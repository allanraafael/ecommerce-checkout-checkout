package queue

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)


func Connect() *amqp.Channel {
	dns := os.Getenv("RABBITMQ_CONNECT_URL")

	conn, err := amqp.Dial(dns)
	if err != nil {
		panic(err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	return channel
}
