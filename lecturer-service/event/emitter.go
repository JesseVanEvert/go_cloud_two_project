package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connection *amqp.Connection
	channel	*amqp.Channel
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()

	if err != nil {
		return err
	}

	channel.QueueDeclare(
		"Messages",    // name?
		true, // durable?
		false, // delete when unused?
		false,  // exclusive?
		false, // no-wait?
		nil,   // arguments?
	)

	e.channel = channel;

	//defer channel.Close()
	return declareExchange(channel)
}

func (e *Emitter) Push(event string, severity string) error {
	log.Println("Pushing to channel")

	err := e.channel.Publish(
		"",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(event),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}
