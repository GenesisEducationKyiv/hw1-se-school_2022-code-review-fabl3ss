package main

import (
	"rabbitmq-log/src/rabbit_mq"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	consumer := rabbit_mq.NewRabbitMQConsumer()
	consumer.LogBindingMessages()
}
