package services

import (
	"encoding/json"
	"testing"

	"../models"
	"github.com/stretchr/testify/assert"
)

func TestGetCharacters(t *testing.T) {
	// Arrange
	var expectedTotalType int16
	var expectedCharactersType []models.Character
	var expectedLinkType models.Link

	// Atc
	charactersResponse, charsErr := GetCharacters()

	var response models.CharactersResponse
	json.Unmarshal(charactersResponse, &response)

	// Asset
	assert.IsType(t, expectedTotalType, response.Total, "The Total property has a wrong type.")
	assert.IsType(t, expectedCharactersType, response.Characters, "The Characters property has a wrong type.")
	assert.IsType(t, expectedLinkType, response.Link, "The Link property has a wrong type.")

	assert.Equal(t, int16(3), response.Total, "Expected to get 3 characters but, %v were retrieved instead.", response.Total)
	assert.NotEmpty(t, response.Characters, "The list of mocked characters should not be empty.")

	assert.Nil(t, charsErr)
}
