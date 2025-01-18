package config

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

/*
Publisher hold the connection to
rabbitmq instance
*/
type Publisher struct {
	conn *amqp091.Connection
}

/*
PublisherChannel hold channel and its
queue on every new channel creation
*/
type PublisherChannel struct {
	channel *amqp091.Channel
	queue   amqp091.Queue
}

/*
NewPublisher function is use to create
new connection
*/
func NewPublisher(url string) *Publisher {

	conn, err := amqp091.Dial(url)

	if err != nil {
		log.Panicf("error when creating consumer connection: %v", err)
	}

	defer conn.Close()

	return &Publisher{
		conn: conn,
	}
}

func (p *Publisher) createChannel() *PublisherChannel {
	ch, err := p.conn.Channel()

	if err != nil {
		log.Fatalf("error when creating consumer channel: %v", err)
	}

	queue, err := ch.QueueDeclare(
		"file-processing-queue", // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)

	if err != nil {
		log.Fatalf("error when creating queue in channel: %v", err)
	}

	return &PublisherChannel{
		channel: ch,
		queue:   queue,
	}
}

func (ch *PublisherChannel) sendMessage(msg []byte) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.channel.PublishWithContext(ctx,
		"",
		ch.queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	if err != nil {
		log.Fatalf("error when sending msg : %v", err)
	}

	log.Printf(" [x] Sent %vn", msg)
}
