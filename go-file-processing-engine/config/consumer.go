package config

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

/*
Consumer hold the connection to
rabbitmq instance
*/
type Consumer struct {
	conn *amqp091.Connection
}

/*
ConsumerChannel hold channel and its
queue on every new channel creation
*/
type ConsumerChannel struct {
	channel *amqp091.Channel
	queue   amqp091.Queue
}

/*
NewPublisher function is use to create
new connection
*/
func NewConsumer(url string) *Consumer {

	conn, err := amqp091.Dial(url)

	if err != nil {
		log.Panicf("error when creating consumer connection: %v", err)
	}

	// defer conn.Close()

	return &Consumer{
		conn: conn,
	}
}

// func (p *Consumer) CreateChannel() *ConsumerChannel {
// 	ch, err := p.conn.Channel()

// 	if err != nil {
// 		log.Fatalf("error when creating consumer channel: %v", err)
// 	}

// 	queue, err := ch.QueueDeclare(
// 		"file-processing-queue", // name
// 		false,                   // durable
// 		false,                   // delete when unused
// 		false,                   // exclusive
// 		false,                   // no-wait
// 		nil,                     // arguments
// 	)

// 	if err != nil {
// 		log.Fatalf("error when creating queue in channel: %v", err)
// 	}

// 	return &ConsumerChannel{
// 		channel: ch,
// 		queue:   queue,
// 	}
// }

// func (ch *ConsumerChannel) ReceiveMessage() any {

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	err := ch.channel.ConsumeWithContext(ctx,
// 		ch.queue.Name,

// 		false,
// 		false,
// 		amqp091.Publishing{
// 			ContentType: "text/plain",
// 			Body:        msg,
// 		})

// 	if err != nil {
// 		log.Fatalf("error when sending msg : %v", err)
// 	}

// 	log.Printf(" [x] Received %vn", msg)
// }
