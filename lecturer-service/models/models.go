package models

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

type ClassLecturerPayload struct {
	ClassId    int `json:"class_id"`
	LecturerId int `json:"lecturer_id"`
}

type IDPayload struct {
	ID int `json:"id"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}