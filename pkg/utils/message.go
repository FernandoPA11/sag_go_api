package utils

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	// The message to be displayed
	Message string
	// The status of the message
	Status string
	// The data to be displayed
	Data interface{}
}

func Respond(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := Message{
		Message: message,
		Status:  http.StatusText(status),
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}
