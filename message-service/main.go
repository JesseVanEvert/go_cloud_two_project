package main

import (
	"fmt"
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

	router.HandleFunc("/messages/lecturer/{lecturerId}", messageHandlers.FindMessageByLecturerId).
		Methods("GET").
		Name(" Message")

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, router))

}
