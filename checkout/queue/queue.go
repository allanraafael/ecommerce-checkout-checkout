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


func Publisher(payload []byte, exchange string, key string, ch *amqp.Channel) {
	err := ch.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing {
			ContentType: "application/json",
			Body: []byte(payload),
		})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Mensagem publicada: " + string(payload))
}
