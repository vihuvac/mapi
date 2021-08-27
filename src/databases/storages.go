package databases

import (
	"../models"
)

type CharacterStorage interface {
	Create(character *models.Character) (*models.CreatedRecordResponse, error)
	FindAll() ([]models.Character, error)
}
