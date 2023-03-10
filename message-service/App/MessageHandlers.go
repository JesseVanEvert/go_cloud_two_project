package app

import (
	Services "MesseageMicroService/restApi/Services"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type MessageHandlers struct {
	Service Services.MessageService
}

func (ch *MessageHandlers) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := ch.Service.GetAllMessages()
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (ch *MessageHandlers) FindByMessageId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message, err := ch.Service.FindMessageById(params["id"])
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, message)
	}
}

// Find message By lecturer email
func (ch *MessageHandlers) FindMessageByLecturerEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message, err := ch.Service.FindMessageByLecturerEmail(params["lecturerEmail"])
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, message)
	}
}
