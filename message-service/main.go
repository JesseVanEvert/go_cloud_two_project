package main

import (
	"fmt"
	"github.com/go-chi/cors"
	"log"
	"net/http"

	app "MesseageMicroService/restApi/App"
	"MesseageMicroService/restApi/Domain"
	"MesseageMicroService/restApi/Services"
	"github.com/gorilla/mux"
)

const webPort = ":8080"

func main() {
	fmt.Println("Starting App")

	var router = mux.NewRouter()

	messageRepo := Domain.NewMessageRepositoryDB()
	messageServices := Services.NewMessageService(messageRepo)

	var messageHandlers = app.MessageHandlers{messageServices}

	router.HandleFunc("/messages", messageHandlers.GetAllMessages).
		Methods("GET").
		Name("GetAllMessages")

	router.HandleFunc("/messages/{id}", messageHandlers.FindByMessageId).
		Methods("GET").
		Name(" Message")

	router.HandleFunc("/messages/lecturer/{lecturerEmail}", messageHandlers.FindMessageByLecturerEmail).
		Methods("GET").
		Name(" Message")

	fmt.Println("Starting Web Server on port", webPort)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
	})

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, c.Handler(router)))

}
