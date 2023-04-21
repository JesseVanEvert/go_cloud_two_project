package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"MessageExchange", 
		"direct",   
		true,       
		false,      
		false,      
		false,      
		nil,        
	)
}

func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"Messages",   
		false, 
		false, 
		true,  
		false, 
		nil,   
	)
}
