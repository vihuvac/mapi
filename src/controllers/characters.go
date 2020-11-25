package controllers

import (
	"net/http"

	"../helpers"
	"../services"
)

// GetCharacters manage the http request.
func GetCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := services.GetCharacters()

	if err == nil && characters == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if charactersError, ok := err.(*helpers.ErrorHandler); ok {
		helpers.ErrorResponse(w, charactersError.Error(), charactersError.Message, charactersError.Code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(characters)
}
