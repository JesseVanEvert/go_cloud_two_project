package models

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type IDPayload struct {
	ID int `json:"id"`
}

type LecturEmailPayload struct {
	Email string `json:"email"`
}