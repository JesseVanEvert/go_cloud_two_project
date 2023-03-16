package event

import (
	"context"
	"encoding/json"
	"fmt"
	"lecturer/ent"
	"lecturer/models"
	"lecturer/repositories"
	Services "lecturer/services"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
	Service  Services.ClassRoomService
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
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

	// connect to database
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/LecturerTest?parseTime=True")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	ctx := context.Background()

	classService := Services.NewClassRoomService(repositories.NewClassRoomRepository(ctx, client))

	consumer.Service = classService

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen() error {
	// Define RabbitMQ server URL.
	amqpServerURL := "amqp://guest:guest@localhost:5672/"
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Subscribing to QueueService1 for getting messages.
	classrooms, err := channelRabbitMQ.Consume(
		"Classes", // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	forever := make(chan bool)
	go func() {
		for classroom := range classrooms {
			var classRoomQueueMessage models.ClassRoomQueueMessage
			_ = json.Unmarshal(classroom.Body, &classRoomQueueMessage)

			go consumer.handleClassRoomMessage(classRoomQueueMessage)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [messages, %s]\n", "Messages")
	<-forever

	return nil
}

func (consumer *Consumer) handleClassRoomMessage(classroom models.ClassRoomQueueMessage) {
	switch classroom.Operation {
	case "DELETE":
		// Delete the classroom
		message, err := consumer.Service.DeleteClassRoom(classroom.ClassRoom.ID)
		if err != nil {
			log.Println(err)
		}
		log.Println(message)
	case "CREATE":
		// Create the classroom
		class, err := consumer.Service.CreateClassRoom(classroom.ClassRoom)
		if err != nil {
			log.Println(err)
		}
		log.Println(class.Name)
	case "UPDATE":
		// Update the classroom
		class, err := consumer.Service.UpdateClassRoom(classroom.ClassRoom)
		if err != nil {
			log.Println(err)
		}
		log.Println(class.Name)
	default:
		log.Println("Invalid operation")
	}
}

/*func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		// log whatever we get
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}

	case "message":
		println("Message received: " + payload.Data)

	default:
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}
	}
}

func logEvent(entry Payload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}*/
