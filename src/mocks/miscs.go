package mocks

import (
	"../helpers"
	"../models"

	uuid "github.com/satori/go.uuid"
)

// GenerateUUID generate a uuid v4.
func GenerateUUID() uuid.UUID {
	uuid, _ := uuid.NewV4()

	return uuid
}

// GenerateItem generate an item.
func GenerateItem(itemProps models.Item) models.Item {
	var item models.Item
	helpers.CopyIdenticalFields(itemProps, &item)

	return item
}

// GenerateCharacter generate a character.
func GenerateCharacter(characterProps models.Character) models.Character {
	var character models.Character
	helpers.CopyIdenticalFields(characterProps, &character)

	return character
}
