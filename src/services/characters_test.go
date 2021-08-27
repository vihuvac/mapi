package services

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"../mocks"
	"../models"
)

type MockCharacterStorage struct {
	mock.Mock
}

func (mock *MockCharacterStorage) Create(character *models.Character) (*models.CreatedRecordResponse, error) {
	args := mock.Called()
	response := args.Get(0)
	return response.(*models.CreatedRecordResponse), args.Error(1)
}

func (mock *MockCharacterStorage) FindAll() ([]models.Character, error) {
	args := mock.Called()
	response := args.Get(0)
	fmt.Printf("|---> CHARACTERS - TEST RESPONSE: %v", reflect.TypeOf(response.([]models.Character)))
	return response.([]models.Character), args.Error(1)
}

func TestGetCharacters(t *testing.T) {
	// Arrange
	mockRepo := new(MockCharacterStorage)

	//var expectedTotalType int16
	var expectedCharactersType []models.Character
	//var expectedLinkType models.Link

	characters := mocks.GenerateCharacters()

	mockRepo.On("FindAll").Return(characters, nil)

	testService := NewCharacterService(mockRepo)

	response, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.IsType(t, expectedCharactersType, response, "The Characters property has a wrong type.")

	/*
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
	*/
}
