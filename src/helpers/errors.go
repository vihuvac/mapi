package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorHandler is the custom error struct for the handlers.
type ErrorHandler struct {
	Message string
	Code    int
}

func (errorHandler *ErrorHandler) Error() string {
	return errorHandler.Message
}

// ControllerError is the custom error struct for the handlers.
type ControllerError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"status"`
}

// ErrorResponse is used by controllers to respond to the API request.
func ErrorResponse(w http.ResponseWriter, err string, message string, code int) {
	log.Printf("[ControllerError]: %s\n", err)

	object := ControllerError{
		Error:   err,
		Message: message,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if j, err := json.Marshal(object); err == nil {
		w.Write(j)
	}
}
