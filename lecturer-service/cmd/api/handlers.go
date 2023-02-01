package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"lecturer/ent"
	"lecturer/event"

	"net/http"
)

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

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}


// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) handleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "message":
		app.putMessageOnQueue(w, requestPayload.Message)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) logItem(w http.ResponseWriter, entry LogPayload) {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)

}

func (app *Config) putMessageOnQueue(w http.ResponseWriter, msg MessagePayload) {

	err := app.pushToQueue(msg.From, msg.To, msg.Message)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Send message via RabbitMQ"

	app.writeJSON(w, http.StatusAccepted, payload)
}

// Change to message instead of log
// pushToQueue pushes a message into RabbitMQ
func (app *Config) pushToQueue(from, to, message string) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
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

func (c *Config) createLecturer(w http.ResponseWriter, lect *http.Request) {
	var lecturerPayload LecturerPayload
	error2 := c.readJSON(w, lect, &lecturerPayload)

	if error2 != nil {
		c.errorJSON(w, error2)
		return 
	}

	var lect2 *ent.Lecturer = &ent.Lecturer{
		FirstName: lecturerPayload.FirstName,
		LastName: lecturerPayload.LastName,
		Email: lecturerPayload.Email,
	}

	lect3, error3 := c.LecturerService.CreateLecturer(lect2)

	var payload jsonResponse

	if error3 != nil {
		payload.Error = true
		payload.Message = "Error creating lecturer"
		payload.Data = error3
		c.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = "Created lecturer"
	payload.Data = lect3

	c.writeJSON(w, http.StatusAccepted, payload)
}

func addLecturerToClass(ctx context.Context, client *ent.Client, lecturerID, classID int) error {
	_, err := client.Class.
		UpdateOneID(classID).
		AddClassLecturerIDs(lecturerID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("adding lecturer to class: %w", err)
	}
	return nil
}

func (app *Config) getAllLecturers(w http.ResponseWriter, lect *http.Request)  {

	ctx :=  context.Background()
	client := ent.Client{}

	lecturers, err := client.Lecturer.
		Query().
		All(ctx)
	if err != nil {
		fmt.Println("getting all lecturers: %w", err)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(lecturers)
}

func getAllClasses(ctx context.Context, client *ent.Client) ([]*ent.Class, error) {
	classes, err := client.Class.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting all classes: %w", err)
	}
	return classes, nil
}

