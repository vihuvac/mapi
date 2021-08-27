package controllers

import (
	"encoding/json"
	"net/http"

	"../helpers"
	"../models"
	"../services"
)

var (
	characterService services.CharacterService
)

type CharacterController interface {
	GetCharacters(w http.ResponseWriter, r *http.Request)
	CreateCharacter(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

func NewCharacterController(service services.CharacterService) CharacterController {
	characterService = service
	return &controller{}
}

// GetCharacters is the controller function to get a list of characters.
func (*controller) GetCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := characterService.FindAll()

	if err == nil && characters == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if charactersError, ok := err.(*helpers.ErrorHandler); ok {
		helpers.ResponseError(w, charactersError.Error(), charactersError.Message, charactersError.Code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(characters)
}

// CreateCharacter is the controller function to create a new character.
func (*controller) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var characterProps models.Character
	err := json.NewDecoder(r.Body).Decode(&characterProps)
	if err != nil {
		helpers.ResponseError(w, err.Error(), "Bad request", http.StatusBadRequest)
		return
	}

	character, err := characterService.Create(characterProps)
	if err != nil {
		if createCharacterError, ok := err.(*helpers.ErrorHandler); ok {
			helpers.ResponseError(w, createCharacterError.Detail, createCharacterError.Message, http.StatusConflict)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(character)
}
