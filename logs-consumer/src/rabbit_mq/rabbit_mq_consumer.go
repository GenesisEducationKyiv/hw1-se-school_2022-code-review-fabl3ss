package rabbit_mq

import (
	"fmt"
	"log"
	"os"
	"rabbitmq-log/src/config/rabbit_mq"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Exchange string
}

type RabbitMQConsumer struct {
	config     rabbitMQConfig
	channel    *amqp.Channel
	connection *amqp.Connection
	queueName  string
	bindingKey string
}

func NewRabbitMQConsumer() *RabbitMQConsumer {
	cfg := buildEnvRabbitMQConfig()
	connection := dialWithRabbitMQ(&cfg)
	channel := configureRabbitMQChannel(&cfg, connection)

	consumer := &RabbitMQConsumer{
		config:     cfg,
		channel:    channel,
		connection: connection,
		queueName:  os.Getenv(config.EnvQueueName),
		bindingKey: os.Getenv(config.EnvBindingKey),
	}
	consumer.configureQueue()

	return consumer
}

func (r *RabbitMQConsumer) LogBindingMessages() {
	messages := r.consumeMessages()
	r.logMessages(messages)
}

func buildEnvRabbitMQConfig() rabbitMQConfig {
	return rabbitMQConfig{
		Host:     os.Getenv(config.EnvRabbitMQHost),
		Port:     os.Getenv(config.EnvRabbitMQPort),
		Username: os.Getenv(config.EnvRabbitMQUser),
		Password: os.Getenv(config.EnvRabbitMQPassword),
		Exchange: os.Getenv(config.EnvRabbitMQExchange),
	}
}

func dialWithRabbitMQ(cfg *rabbitMQConfig) *amqp.Connection {
	rabbitURL := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
	connection, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatal(err, "unable connect to RabbitMQ")
		return nil
	}

	return connection
}

func configureRabbitMQChannel(cfg *rabbitMQConfig, conn *amqp.Connection) *amqp.Channel {
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalln(err.Error(), "unable to create channel")
	}

	if err := channel.ExchangeDeclare(
		cfg.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Fatalln(err.Error(), "failed to declare exchange")
	}

	return channel
}

func (r *RabbitMQConsumer) configureQueue() {
	if _, err := r.channel.QueueDeclare(
		r.queueName,
		false,
		false,
		true,
		false,
		nil,
	); err != nil {
		log.Fatalln(err.Error(), "unable to declare queue")
	}

	if err := r.channel.QueueBind(
		r.queueName,
		r.bindingKey,
		r.config.Exchange,
		false,
		nil,
	); err != nil {
		log.Fatalln(err.Error(), "unable to bind queue")
	}
}

func (r *RabbitMQConsumer) consumeMessages() <-chan amqp.Delivery {
	messages, err := r.channel.Consume(
		r.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalln(err.Error(), "unable to consume channel")
	}

	return messages
}

func (r *RabbitMQConsumer) logMessages(messages <-chan amqp.Delivery) {
	var forever chan struct{}

	go func() {
		for d := range messages {
			log.Printf("%s", d.Body)
		}
	}()

	log.Printf("Start listening log queue")
	<-forever
}
