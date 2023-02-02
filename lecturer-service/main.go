package main

import (
	"context"
	"encoding/json"
	"fmt"
	"lecturer/ent"
	"lecturer/event"
	"lecturer/helpers"
	models "lecturer/models"
	"lecturer/repositories"
	Services "lecturer/services"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "80"


type Config struct {
	Rabbit *amqp.Connection
	Helpers helpers.Helpers
	Service Services.LecturerService
}


func (c *Config) registerRoutes() {
	mux := chi.NewRouter()

	mux.HandleFunc("/createLecturer", c.CreateLecturer)
	mux.HandleFunc("/getAllLecturers", c.GetAllLecturers)
	mux.HandleFunc("/addLecturerToClass", c.AddLecturerToClass)
	mux.HandleFunc("/sendMessage", c.SendMessage)
	mux.HandleFunc("/getAllClasses", c.GetAllClasses)
	mux.HandleFunc("/getLecturerByID", c.GetLecturerByID)

	http.ListenAndServe(":8080", mux)
}
func main() {


	// connect to database
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/lecturerTest?parseTime=True")

	if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
	
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()

	log.Printf("Starting broker service on port %s\n", webPort)

	lecturerService := Services.NewLecturerService(repositories.NewLecturerRepository(ctx, client))

	c := Config{
		Rabbit: rabbitConn,
		Helpers: helpers.NewHelpers(),
		Service: lecturerService,
	}

	c.registerRoutes()
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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


// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (c *Config ) SendMessage(w http.ResponseWriter, r *http.Request) {
	var requestPayload models.RequestPayload
	
	err := c.Helpers.ReadJSON(w, r, &requestPayload)
	
	//err := h.h.readJSON(w, r, &requestPayload)
	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "message":
		c.putMessageOnQueue(w, requestPayload.Message)
	default:
		c.Helpers.ErrorJSON(w, err)
	}
}

func (c *Config )  putMessageOnQueue(w http.ResponseWriter, msg models.MessagePayload) {

	err := c.pushToQueue(msg.From, msg.To, msg.Message)
	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Send message via RabbitMQ"

	c.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

// Change to message instead of log
// pushToQueue pushes a message into RabbitMQ
func (c *Config)  pushToQueue(from, to, message string) error {
	emitter, err := event.NewEventEmitter(c.Rabbit)
	if err != nil {
		return err
	}

	payload := models.MessagePayload{
		From:    from,
		To:      to,
		Message: message,
	}

	j, _ := json.MarshalIndent(&payload, "", "\t")
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		return err
	}
	return nil
}

func (c *Config ) CreateLecturer(w http.ResponseWriter, lect *http.Request) {
	var lecturerPayload models.LecturerPayload
	error := c.Helpers.ReadJSON(w, lect, &lecturerPayload)

	if error!= nil {
		c.Helpers.ErrorJSON(w, error)
		return 
	}

	lect3, error := c.Service.CreateLecturer(lecturerPayload)

	if(error != nil){
		c.Helpers.ErrorJSON(w, error)
		return
	}

	var payload models.JsonResponse

	payload.Error = false
	payload.Message = "Created lecturer"
	payload.Data = lect3

	c.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (c *Config) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	classes, error := c.Service.GetAllClasses()

	if(error != nil){
		c.Helpers.ErrorJSON(w, error)
		return
	}

	var payload models.JsonResponse

	payload.Error = false
	payload.Message = "All classes"
	payload.Data = classes

	c.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (c *Config ) AddLecturerToClass(w http.ResponseWriter, lect *http.Request) {
	var classLecturerPayload models.ClassLecturerPayload
	error := c.Helpers.ReadJSON(w, lect, &classLecturerPayload)

	if(error != nil){
		c.Helpers.ErrorJSON(w, error)
		return
	}

	message, error := c.Service.AddLecturerToClass(classLecturerPayload.ClassId, classLecturerPayload.LecturerId)

	if(error != nil){
		c.Helpers.ErrorJSON(w, error)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = message
	payload.Data = nil

	c.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (c *Config ) GetAllLecturers(w http.ResponseWriter, request *http.Request)  {
	lecturers, err := c.Service.GetAllLecturers()

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Retrieved lecturers"
	payload.Data = lecturers

	c.Helpers.WriteJSON(w, http.StatusOK, payload)
}

func (c *Config) GetLecturerByID (w http.ResponseWriter, request *http.Request) {
	var idPayload models.IDPayload
	err := c.Helpers.ReadJSON(w, request, &idPayload)

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	lecturer, err := c.Service.GetLecturerByID(idPayload.ID)

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Retrieved lecturer"
	payload.Data = lecturer

	c.Helpers.WriteJSON(w, http.StatusOK, payload)
}




