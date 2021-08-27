package databases

import (
	"net/http"

	"github.com/lib/pq" // Lib included by gorm.

	"../helpers"
	"../models"
	"./sqlite"
)

type sqliteStruct struct{}

func NewSQLiteStorage() CharacterStorage {
	return &sqliteStruct{}
}

func (sql *sqliteStruct) Create(character *models.Character) (*models.CreatedRecordResponse, error) {
	// Create the new character record.
	if err := sqlite.DB.Create(&character).Error; err != nil {
		// Get the properties from the error type *pq.Error retrieved from the DB exception.
		pqErr := err.(*pq.Error)
		return nil, &helpers.ErrorHandler{pqErr.Detail, "There was an error trying to create a new character.", http.StatusConflict}
	}

	response := &models.CreatedRecordResponse{character.ID}
	return response, nil
}

func (sql *sqliteStruct) FindAll() ([]models.Character, error) {
	characters := make([]models.Character, 0)
	sqlite.DB.Find(&characters)
	return characters, nil
}
