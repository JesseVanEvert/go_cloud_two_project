package event

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	services "messages/services"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MessagePayload struct {
	From    string `json:"from"`
	Message string `json:"message"`
	To      []string `json:"to"`
}

type MessageDB struct {
	db *sql.DB
}

type Consumer struct {
	conn      *amqp.Connection
	Service  services.MessageService
}

func NewConsumer(conn *amqp.Connection, service services.MessageService) (Consumer, error) {
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

	messages, err := channel.Consume(
		"Messages", 
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
			for _, email := range payload.To {
				log.Println(email)
				ch.InsertMessage(payload.From, payload.Message, email)
				log.Println("insert Successful")
			}
		}
	}()
	<-forever

	return nil
}


func postMessage(payload MessagePayload) error {
	for _, email := range payload.To {
		content := payload.Message
		from := mail.NewEmail("User", payload.From)
		subject := "Email from teacher"
		to := mail.NewEmail("Example User", email)
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
