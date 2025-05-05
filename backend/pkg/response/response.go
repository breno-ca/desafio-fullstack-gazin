package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

const encodingResponseError string = "error encoding response data"

func Write(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.Write([]byte(encodingResponseError))
		return
	}
}

func WriteError(w http.ResponseWriter, statusCode int, message string, error error) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(ErrorResponse{
		Error:   error.Error(),
		Message: message,
	})
	if err != nil {
		w.Write([]byte(encodingResponseError))
		return
	}
}
