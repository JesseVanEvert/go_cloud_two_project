package event

import (
	"encoding/json"
	"fmt"
	"lecturer/models"
	Services "lecturer/services"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	Service  Services.ClassRoomService
}

func NewConsumer(conn *amqp.Connection, service Services.ClassRoomService) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
		Service: service,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen() error {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channel, err := consumer.conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	declareQueue(channel)

	classrooms, err := channel.Consume(
		"Classes", 
		"",        
		true,      
		false,     
		false,     
		false,     
		nil,       
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	forever := make(chan bool)
	go func() {
		for classroom := range classrooms {
			log.Println(classroom)
			var classRoomQueueMessage models.ClassRoomQueueMessage
			_ = json.Unmarshal(classroom.Body, &classRoomQueueMessage)
			go consumer.handleClassRoomMessage(classRoomQueueMessage)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [MessageExchange, %s]\n", "Messages")
	<-forever

	return nil
}

func (consumer *Consumer) handleClassRoomMessage(classroomMessage models.ClassRoomQueueMessage) {
	var classroom =  models.ClassRoom{ID: classroomMessage.ClassRoomId, Classname: classroomMessage.ClassRoom}
	switch classroomMessage.Operation {
	case "DELETE":
		// Delete the classroom
		message, err := consumer.Service.DeleteClassRoom(classroom.ID)
		if err != nil {
			log.Println(err)
		}
		log.Println(message)
	case "CREATE":
		// Create the classroom
		class, err := consumer.Service.CreateClassRoom(classroom)
		if err != nil {
			log.Println(err)
		}
		log.Println(class.Name)
	case "UPDATE":
		// Update the classroom
		class, err := consumer.Service.UpdateClassRoom(classroom)
		if err != nil {
			log.Println(err)
		}
		log.Println(class.Name)
	default:
		log.Println("Invalid operation")
	}
}

