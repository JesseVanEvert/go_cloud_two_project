package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/streadway/amqp"
)

type MessagePayload struct {
	Message string `json:"message"`
	From    string `json:"from"`
	To      []Recipient
}

type Recipient struct {
	Email string `json:"to"`
}

type MessageDB struct {
	db *sql.DB
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
			ch := &MessageDB{db: NewMessageRepositoryDB().db}
			var payload MessagePayload
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				fmt.Println("Error parsing message payload:", err)
			}
			postMessage(payload)
			log.Printf("Received a message: %s", d.Body)
			for _, recipient := range payload.To {
				log.Println(recipient.Email)
				ch.InsertMessage(payload.From, payload.Message, recipient.Email)
				log.Println("insert Successful")
			}
		}
	}()
	<-forever
}

func postMessage(payload MessagePayload) error {
	for _, recipient := range payload.To {
		content := payload.Message
		from := mail.NewEmail("User", payload.From)
		subject := "Email from teacher"
		to := mail.NewEmail("Example User", recipient.Email)
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
	}
	return nil
}

func (ch *MessageDB) InsertMessage(lectureEmail string, msg string, toEmail string) (int64, error) {
	res, err := ch.db.Exec("INSERT INTO message (lecturerEmail , content, toEmail) VALUES (?,?,?)", lectureEmail, msg, toEmail)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	log.Println("inserting is working", id)
	return id, nil
}

func NewMessageRepositoryDB() MessageDB {

	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/message")

	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return MessageDB{db}
}
