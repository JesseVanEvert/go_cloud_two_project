package main

import (
	"context"
	"fmt"
	"lecturer/ent"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	ls "lecturer/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "80"


type Config struct {
	Rabbit *amqp.Connection
	LecturerService ls.LecturerService
}

func (c *Config) registerRoutes() {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	//mux.Post("/", app.Broker)

	mux.Post("/handle", c.handleSubmission)
	mux.Post("/createLecturer", c.createLecturer)
	//mux.Post("/handle", app.HandleSubmission)
	//mux.Post("/createLecturer", app.CreateLecturer)
	//mux.Post("/getAllLecturers", app.GetAllLecturers)
	http.ListenAndServe(":8080", mux)
}
func main() {

	// connect to database
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/lecturer?parseTime=True")

	if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
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

	c := Config{
		Rabbit: rabbitConn,
	}

	c.registerRoutes()

	// define http server
	/*srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.registerRoutes(),
	}*/

	// start the server
	/*err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}*/
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


