package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"MessageExchange", // name
		"direct",    // type
		true,       // durable?
		false,      // auto-deleted?
		false,      // internal?
		false,      // no-wait?
		nil,        // arguements?
	)
}

func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"Messages",    // name?
		true, // durable?
		false, // delete when unused?
		true,  // exclusive?
		false, // no-wait?
		nil,   // arguments?
	)
}