package config

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

/*
Consumer hold the connection to
rabbitmq instance
*/
type Consumer struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   amqp091.Queue
}

/*
NewConsumer function is use to create
new connection
*/
func NewConsumer(url string, queueName string) *Consumer {

	conn, err := amqp091.Dial(url)

	if err != nil {
		log.Panicf("error when creating consumer connection: %v", err)
	}

	// defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("error when creating consumer channel: %v", err)
	}

	queue, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		log.Fatalf("error when creating queue in channel: %v", err)
	}

	return &Consumer{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}
}

func (c *Consumer) ReceiveMessage() (<-chan amqp091.Delivery, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgs, err := c.channel.ConsumeWithContext(
		ctx,
		c.queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("error when reading msg from broker : %v", err)
		return msgs, err
	}

	return msgs, nil
}
