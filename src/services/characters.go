package services

import (
	"encoding/json"

	"../databases/postgres"
	"../helpers"
	"../mocks"
	"../models"

	validator "gopkg.in/go-playground/validator.v9"
)

// GetCharacters is the service to manage the characters logic.
func GetCharacters() ([]byte, error) {
	// Instantiate the mocks generator.
	characters := mocks.GenerateCharacters()

	// Create the API response structure.
	apiResponse := &models.CharactersResponse{Total: int16(len(characters)), Characters: characters, Link: models.Link{Next: ""}}

	// Convert the response into bytes.
	response, errMarshal := json.Marshal(apiResponse)
	if errMarshal != nil {
		return nil, &helpers.ErrorHandler{"Something went trying to get characters.", 500}
	}

	return response, nil
}

// CreateCharacter service used to create a character doing some validations.
func CreateCharacter(characterProps models.Character) ([]byte, error) {
	// Validate the character props.
	validate := validator.New()
	errValidation := validate.Struct(characterProps)
	if errValidation != nil {
		return nil, &helpers.ErrorHandler{"Bad request", 400}
	}

	// Create the new character record.
	var character models.Character
	helpers.CopyIdenticalFields(characterProps, &character)
	postgres.DB.Create(&character)

	// Create the response.
	response, err := json.Marshal(character)
	if err != nil {
		return nil, &helpers.ErrorHandler{"Server error", 500}
	}

	return response, nil
}
