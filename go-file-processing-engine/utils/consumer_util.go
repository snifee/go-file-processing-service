package utils

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

var (
	url = viper.GetString("rabbitmq.url")
)

type Consumer struct {
	conn *amqp091.Connection
}

func NewConsumer() *Consumer {
	conn, err := amqp091.Dial(url)

	if err != nil {
		log.Panicf("error when creating consumer connection: %v", err)
	}

	defer conn.Close()

	return &Consumer{
		conn: conn,
	}
}
