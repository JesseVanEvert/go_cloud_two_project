package main

import (
	/*"fmt"
	"log"
	"net/http"

	"github.com/go-chi/cors"

	app "MesseageMicroService/restApi/App"
	"MesseageMicroService/restApi/Domain"
	"MesseageMicroService/restApi/Services"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/gorilla/mux"*/

	"log"
	"net/http"
	"os"
	models "messages/models"
	"messages/repositories"
	services "messages/services"
	"math"
	"time"

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
	messageRepo := domain.NewMessageRepositoryDB()
	messageServices := services.NewMessageService(messageRepo)


	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()

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
	/*params := mux.Vars(r)
	message, err := c.MessageService.FindMessageById(params["id"])

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, message)
	}*/

}

// Find message By lecturer email
func (c *Config) FindMessageByLecturerEmail(w http.ResponseWriter, request *http.Request) {
	var lecturerEmailPayload models.LecturerEmailPayload
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

	/*if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, message)
	}*/
}
