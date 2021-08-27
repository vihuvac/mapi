package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorHandler is the custom error struct for the handlers.
type ErrorHandler struct {
	Detail  string
	Message string
	Code    int
}

func (errorHandler *ErrorHandler) Error() string {
	return errorHandler.Message
}

// ErrorController is the custom error struct for the handlers.
type ErrorController struct {
	Detail  string `json:"detail"`
	Message string `json:"message"`
}

// ResponseError is used by controllers to respond to the API request.
func ResponseError(w http.ResponseWriter, detail string, message string, code int) {
	log.Printf("[Response Error]: %s\n", detail)

	object := ErrorController{
		Detail:  detail,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if j, err := json.Marshal(object); err == nil {
		w.Write(j)
	}
}
