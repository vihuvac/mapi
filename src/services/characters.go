package services

import (
	"encoding/json"

	"../helpers"
	"../mocks"
	"../models"
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
