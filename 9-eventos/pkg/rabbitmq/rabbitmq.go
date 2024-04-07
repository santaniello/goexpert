package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	// Open a channel
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, channelOut chan<- amqp.Delivery, queue string) error {
	// Consume messages
	msgs, err := ch.Consume(
		queue,         // queue
		"go-consumer", // consumer
		false,         // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return err
	}

	// Get messages from the channel
	for msg := range msgs {
		channelOut <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, body string, exName string) error {
	err := ch.PublishWithContext(
		context.Background(),
		exName, // exchange default from rabbitmq
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
