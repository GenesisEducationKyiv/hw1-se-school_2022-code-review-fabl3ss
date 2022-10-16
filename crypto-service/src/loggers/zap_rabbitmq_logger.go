package loggers

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"genesis_test_case/src/config"
	"io"
	"log"
	"net/url"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

const (
	debugLogKey       = "debug"
	infoLogKey        = "info"
	errorLogKey       = "error"
	rabbitmqWriterKey = "rabbitmqwritter"
)

type customWriter struct {
	io.Writer
}

func (cw customWriter) Close() error {
	return nil
}

func (cw customWriter) Sync() error {
	return nil
}

type rabbitMQConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Exchange string
}

type zapRabbitMQLogger struct {
	buffer     *bytes.Buffer
	writer     *bufio.Writer
	logger     *zapLogger
	channel    *amqp.Channel
	connection *amqp.Connection
	config     rabbitMQConfig
}

func NewZapRabbitMQLogger() *zapRabbitMQLogger {
	logPath := fmt.Sprintf("%s:", rabbitmqWriterKey)
	cfg := buildEnvRabbitMQConfig()
	connection := dialWithRabbitMQ(&cfg)
	channel := configureRabbitMQChannel(&cfg, connection)
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	regLoggerSink(writer)
	logger := &zapRabbitMQLogger{
		buffer:     &buffer,
		writer:     writer,
		logger:     NewZapLogger(logPath),
		config:     cfg,
		channel:    channel,
		connection: connection,
	}

	return logger
}

func (l *zapRabbitMQLogger) Debug(msg string) {
	l.logger.Debug(msg)
	l.publishLog(debugLogKey)
}

func (l *zapRabbitMQLogger) Info(msg string) {
	l.logger.Info(msg)
	l.publishLog(infoLogKey)
}

func (l *zapRabbitMQLogger) Error(msg string) {
	l.logger.Error(msg)
	l.publishLog(errorLogKey)
}

func (l *zapRabbitMQLogger) Close() {
	err := l.channel.Close()
	log.Println(err.Error(), "failed to close a channel")

	err = l.connection.Close()
	log.Println(err.Error(), "failed to close connection to RabbitMQ")
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

func configureRabbitMQChannel(cfg *rabbitMQConfig, connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err, "unable to configure RabbitMQ channel")
		return nil
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
		log.Fatal(err, "unable to declare exchange")
	}
	return channel
}

func regLoggerSink(writer *bufio.Writer) {
	err := zap.RegisterSink(rabbitmqWriterKey, func(u *url.URL) (zap.Sink, error) {
		return customWriter{writer}, nil
	})
	if err != nil {
		log.Fatal(err, "unable to sink logger")
	}
}

func (l *zapRabbitMQLogger) publishLog(key string) {
	_ = l.writer.Flush()
	err := l.channel.PublishWithContext(context.Background(),
		l.config.Exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        l.buffer.Bytes(),
		},
	)
	if err != nil {
		log.Fatal(err.Error(), "failed to publish a message")
	}
	l.buffer.Reset()
}
