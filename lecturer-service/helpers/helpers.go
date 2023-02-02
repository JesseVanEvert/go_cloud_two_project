package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Helpers interface{
	ReadJSON(w http.ResponseWriter, r *http.Request, data any) error
	WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error
	ErrorJSON(w http.ResponseWriter, err error) error
}

type DefaultHelpers struct {
	
}

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

// readJSON tries to read the body of a request and converts it into JSON
func (dh DefaultHelpers) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

// writeJSON takes a response status code and arbitrary data and writes a json response to the client
func (dh DefaultHelpers) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON takes an error, and optionally a response status code, and generates and sends
// a json error response
func (dh DefaultHelpers) ErrorJSON(w http.ResponseWriter, err error) error {
	statusCode := http.StatusBadRequest

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return dh.WriteJSON(w, statusCode, payload)
}

func NewHelpers() DefaultHelpers {
	return DefaultHelpers{}
}