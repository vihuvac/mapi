package services

import (
	"encoding/json"
	"net/http"

	"../databases"
	"../helpers"
	"../models"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	db databases.CharacterStorage
)

type CharacterService interface {
	Create(character models.Character) ([]byte, error)
	FindAll() ([]byte, error)
}

type service struct{}

func NewCharacterService(database databases.CharacterStorage) CharacterService {
	db = database
	return &service{}
}

func validate(character models.Character) error {
	// Validate the character props.
	validate := validator.New()
	errValidation := validate.Struct(character)
	if errValidation != nil {
		return &helpers.ErrorHandler{"Parameters sent are not valid.", "Bad request", http.StatusBadRequest}
	}
	return nil
}

func (s *service) Create(characterProps models.Character) ([]byte, error) {
	err := validate(characterProps)
	if err != nil {
		return nil, err
	}

	// Create the new character record.
	var character models.Character
	helpers.CopyIdenticalFields(characterProps, &character)
	createdCharacter, err := db.Create(&character)
	if err != nil {
		return nil, err
	}

	// Create the response.
	response, err := json.Marshal(createdCharacter)
	if err != nil {
		return nil, &helpers.ErrorHandler{"There was an error encoding the response to JSON.", "Server error", http.StatusInternalServerError}
	}

	return response, nil
}

func (s *service) FindAll() ([]byte, error) {
	characters, err := db.FindAll()
	if err != nil {
		return nil, &helpers.ErrorHandler{"There was an error trying to find records in the DB.", "There was an error trying to retrieve characters from the DB.", http.StatusInternalServerError}
	}

	// Create the API response structure.
	apiResponse := &models.CharactersResponse{Total: int16(len(characters)), Characters: characters, Link: models.Link{Next: ""}}

	// Convert the response into bytes.
	response, errMarshal := json.Marshal(apiResponse)
	if errMarshal != nil {
		return nil, &helpers.ErrorHandler{"There was an error encoding the response to JSON.", "Something went wrong trying to format characters response.", http.StatusInternalServerError}
	}

	return response, nil
}
