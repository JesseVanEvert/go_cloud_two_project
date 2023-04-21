package main

import (
	"context"
	"fmt"
	API "lecturer/controllers"
	"lecturer/ent"
	"lecturer/event"
	"lecturer/helpers"
	"lecturer/repositories"
	"lecturer/services"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

/*

VERANDEREN NAAR ENV

*/

const webPort = "80"

type Config struct {
	Rabbit *amqp.Connection
	Helpers helpers.Helpers
	LecturerService services.LecturerService
	ClassService services.ClassRoomService
	Channel *amqp.Channel
}


func main() {
	// connect to database
	client, err := ent.Open(os.Getenv("lECTURER_DATABASE_TYPE"), os.Getenv("LECTURER_MYSQL_CONNECTION_STRING"))

	log.Println(os.Getenv("lECTURER_DATABASE_TYPE"))
	log.Println(os.Getenv("LECTURER_MYSQL_CONNECTION_STRING"))

	if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }

	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	log.Printf("Starting broker service on port %s\n", webPort)

	lecturerService := services.NewLecturerService(repositories.NewLecturerRepository(ctx, client))
	classService := services.NewClassRoomService(repositories.NewClassRoomRepository(ctx, client))

	/* 

	VERPLAATSEN NAAR EIGEN PACKAGE

	*/

	c := Config{
		Rabbit: rabbitConn,
		Helpers: helpers.NewHelpers(),
		LecturerService: lecturerService,
		ClassService: classService,
	}

	consumer, err := event.NewConsumer(rabbitConn, classService)
	if err != nil {
		log.Println(err)
	}

	go consumer.Listen()

	api := API.NewAPI(c.Rabbit, &c.Helpers, &c.LecturerService, &c.ClassService)

	go api.Start()
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

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







