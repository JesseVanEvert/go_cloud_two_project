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

type Payload struct {
	msg string `json:"message"`
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
			msg := string(d.Body)
			ch := &MessageDB{db: NewMessageRepositoryDB().db}

			/*			if !json.Valid(d.Body) {
						fmt.Println("Error: Invalid JSON message")
						continue
					}*/
			/*			var payload Payload
						err := json.Unmarshal(d.Body, &payload)
						if err != nil {
							fmt.Println("Error parsing message payload:", err)
							continue
						}*/

			// Pass payload to postMessage method
			postMessage(msg)
			log.Printf("Received a message: %s", d.Body)
			ch.InsertMessage(6, msg)
		}
	}()

	<-forever
}

func postMessage(entry string) error {
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

func (ch *MessageDB) InsertMessage(lectureId int, msg string) (int64, error) {
	res, err := ch.db.Exec("INSERT INTO message (lecturerID,content) VALUES (?,?)", lectureId, msg)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
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
