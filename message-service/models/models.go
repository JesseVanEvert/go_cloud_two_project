package models

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type IDPayload struct {
	ID string `json:"id"`
}

type LecturEmailPayload struct {
	Email string `json:"email"`
}