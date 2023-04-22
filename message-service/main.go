package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"messages/event"
	"messages/helpers"
	models "messages/models"
	"messages/repositories"
	services "messages/services"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = ":8008"

type Config struct {
	Rabbit *amqp.Connection
	Helpers helpers.Helpers
	MessageService services.MessageService
	Channel *amqp.Channel
}

func (c *Config) registerRoutes() {
	mux := chi.NewRouter()

	mux.HandleFunc("/messages", c.GetAllMessages)
	mux.HandleFunc("/messages/{id}", c.FindByMessageId)
	mux.HandleFunc("/messages/lecturer/{lecturerEmail}", c.FindMessageByLecturerEmail)

	http.ListenAndServe(":8080", mux)

	/*
	fmt.Println("Starting Web Server on port", webPort)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
	})

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, c.Handler(router)))
	*/
}

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()

	db, err := sql.Open(os.Getenv("MESSAGE_DATABASE_TYPE"), os.Getenv("MESSAGE_MYSQL_CONNECTION_STRING"))

	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	messageRepo := repositories.NewMessageRepository(db)
	messageServices := services.NewMessageService(messageRepo)

	log.Printf("Starting broker service on port %s\n", webPort)

	c := Config{
		Rabbit: rabbitConn,
		Helpers: helpers.NewHelpers(),
		MessageService: messageServices,
	}

	consumer, err := event.NewConsumer(rabbitConn, messageServices)
	if err != nil {
		log.Println(err)
	}

	go consumer.Listen()

	c.registerRoutes()
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}

func (c *Config) GetAllMessages(w http.ResponseWriter, request *http.Request) {
	messages, err := c.MessageService.GetAllMessages()

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Retrieved messages"
	payload.Data = messages

	//w.Header().Set("Content-Type", "application/json")
	c.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (c *Config) FindByMessageId(w http.ResponseWriter, request *http.Request) {
	var idPayload models.IDPayload
	err := c.Helpers.ReadJSON(w, request, &idPayload)

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	message, err := c.MessageService.FindMessageById(idPayload.ID)

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Retrieved message"
	payload.Data = message

	c.Helpers.WriteJSON(w, http.StatusOK, payload)
}

// Find message By lecturer email
func (c *Config) FindMessageByLecturerEmail(w http.ResponseWriter, request *http.Request) {
	var lecturerEmailPayload models.LecturEmailPayload
	err := c.Helpers.ReadJSON(w, request, &lecturerEmailPayload)
	
	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	message, err := c.MessageService.FindMessageByLecturerEmail(lecturerEmailPayload.Email)

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Retrieved message"
	payload.Data = message

	c.Helpers.WriteJSON(w, http.StatusOK, payload)
}
