package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/streadway/amqp"
)

type Payload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func main() {
	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

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
	messages, err := channelRabbitMQ.Consume(
		"QueueService1", // queue name
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

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for d := range messages {
			// For example, show received message in a console.
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-forever
}
func postMessage(entry Payload) error {

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	content := string(jsonData)
	from := mail.NewEmail("User", "mridulhasan157@gmail.com")
	subject := "Email from teacher"
	to := mail.NewEmail("Example User", "mahedimridul57@gmail.com")
	htmlContent := "<strong>Important!</strong>"
	message := mail.NewSingleEmail(from, subject, to, content, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		log.Println("Did not send")
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}
