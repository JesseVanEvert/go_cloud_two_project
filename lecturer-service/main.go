package main

import (
	"context"
	"encoding/json"
	"fmt"
	"lecturer/ent"
	"lecturer/event"
	"lecturer/helpers"
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

type RequestPayload struct {
	Action  string         `json:"action"`
	Auth    AuthPayload    `json:"auth,omitempty"`
	Log     LogPayload     `json:"log,omitempty"`
	Message MessagePayload `json:"message,omitempty"`
}

type MessagePayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LecturerPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}


func (c *Config) registerRoutes() {


	mux := chi.NewRouter()

	// specify who is allowed to connect
	/*mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))*/

	mux.HandleFunc("/createLecturer", c.createLecturer)
	mux.HandleFunc("/getAllLecturers", c.getAllLecturers)
	mux.HandleFunc("/handle", c.handleSubmission)

	http.ListenAndServe(":8080", mux)
}
func main() {


	// connect to database
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/lecturer?parseTime=True")

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
func (c *Config ) handleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	
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

func (c *Config )  putMessageOnQueue(w http.ResponseWriter, msg MessagePayload) {

	err := c.pushToQueue(msg.From, msg.To, msg.Message)
	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	var payload jsonResponse
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

	payload := MessagePayload{
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

func (c *Config ) createLecturer(w http.ResponseWriter, lect *http.Request) {
	var lecturerPayload LecturerPayload
	error2 := c.Helpers.ReadJSON(w, lect, &lecturerPayload)

	if error2 != nil {
		c.Helpers.ErrorJSON(w, error2)
		return 
	}

	var lect2 *ent.Lecturer = &ent.Lecturer{
		FirstName: lecturerPayload.FirstName,
		LastName: lecturerPayload.LastName,
		Email: lecturerPayload.Email,
	}

	lect3, error3 := c.Service.CreateLecturer(lect2)

	var payload jsonResponse

	if error3 != nil {
		payload.Error = true
		payload.Message = "Error creating lecturer"
		payload.Data = error3
		c.Helpers.WriteJSON(w, http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = "Created lecturer"
	payload.Data = lect3

	c.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (c *Config )  addLecturerToClass(ctx context.Context, client *ent.Client, lecturerID, classID int) error {
	_, err := client.Class.
		UpdateOneID(classID).
		AddClassLecturerIDs(lecturerID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("adding lecturer to class: %w", err)
	}
	return nil
}

func (c *Config ) getAllLecturers(w http.ResponseWriter, lect *http.Request)  {
	lecturers, err := c.Service.GetAllLecturers()

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Retrieved lecturers"
	payload.Data = lecturers

	c.Helpers.WriteJSON(w, http.StatusOK, payload)
}

func (c *Config )  getAllClasses(w http.ResponseWriter, lect *http.Request) ([]*ent.Class, error) {
	classes, err := c.Service.GetAllClasses()

	if err != nil {
		c.Helpers.ErrorJSON(w, err)
		return nil, err
	}
	return classes, nil
}



