package databases

import (
	"fmt"
	"net/http"

	"../helpers"
	"../models"
	"./postgres"

	"github.com/lib/pq" // Lib included by gorm.
)

type pgStruct struct{}

func NewPostgresQLStorage() CharacterStorage {
	return &pgStruct{}
}

func (dbs *pgStruct) Create(character *models.Character) (*models.CreatedRecordResponse, error) {
	// Create the new character record.
	if err := postgres.DB.Create(&character).Error; err != nil {
		// Get the properties from the error type *pq.Error retrieved from the DB exception.
		pqErr := err.(*pq.Error)
		return nil, &helpers.ErrorHandler{pqErr.Detail, "There was an error trying to create a new character.", http.StatusConflict}
	}

	response := &models.CreatedRecordResponse{character.ID}
	return response, nil
}

func (dbs *pgStruct) FindAll() ([]models.Character, error) {
	characters := make([]models.Character, 0)
	postgres.DB.Find(&characters)
	fmt.Printf("|---> CHARACTERS: %v", characters)
	return characters, nil
}
