package controllers

import (
	"encoding/json"
	"net/http"

	"../helpers"
	"../models"
	"../services"
)

// GetCharacters is the controller function to get a list of characters.
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

// CreateCharacter is the controller function to create a new character.
func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var characterProps models.Character
	err := json.NewDecoder(r.Body).Decode(&characterProps)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), "Bad request", 400)
		return
	}

	character, err := services.CreateCharacter(characterProps)
	if err != nil {
		if createCharacterError, ok := err.(*helpers.ErrorHandler); ok {
			helpers.ErrorResponse(w, createCharacterError.Error(), createCharacterError.Message, createCharacterError.Code)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(character)
}
