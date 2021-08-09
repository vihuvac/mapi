package mocks

import (
	"time"

	"../models"
)

// GenerateCharacters generate a list of characters used for testing purposes.
func GenerateCharacters() []models.Character {
	// Init characters var as a slice.
	var characters []models.Character

	// Create Harry's items.
	var harryItems []models.Item
	harryItem := models.Item{
		Type:            "Wand",
		Length:          11, // Floating number to represent the length in inches.
		Characteristics: "Nice and supple, brother of Voldemort's wand.",
		Manufacturer:    "Garrick Ollivander",
		Base: models.Base{
			ID:        GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Create Harry's character.
	harry := models.Character{
		FirstName: "Harry",
		LastName:  "Potter",
		Nickname:  "Harry",
		Base: models.Base{
			ID:        GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Items: append(harryItems, harryItem),
	}
	//harry.BeforeCreate() // This method will behave as a hook when the DB is implemented.
	characters = append(characters, harry)

	// Create Ron's items.
	var ronItems []models.Item
	ronItem := models.Item{
		Type:            "Wand",
		Length:          12, // Floating number to represent the length in inches.
		Characteristics: "Made of ash wood, and possessed a unicorn hair core.",
		Manufacturer:    "Garrick Ollivander",
		Base: models.Base{
			ID:        GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Create Ron's character.
	ron := models.Character{
		FirstName: "Ronald",
		LastName:  "Weasley",
		Nickname:  "Ron",
		Base: models.Base{
			ID:        GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Items: append(ronItems, ronItem),
	}
	//ron.BeforeCreate() // This method will behave as a hook when the DB is implemented.
	characters = append(characters, ron)

	// Create Hermione's items.
	var hermioneItems []models.Item
	hermioneItem := models.Item{
		Type:            "Wand",
		Length:          10.75, // Floating number to represent the length in inches.
		Characteristics: "Made of vine wood, and possessed a dragon heartstring core.",
		Manufacturer:    "Garrick Ollivander",
		Base: models.Base{
			ID:        GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Create Hermione's character.
	hermione := models.Character{
		FirstName: "Hermione",
		LastName:  "Granger",
		Nickname:  "Hermione",
		Base: models.Base{
			ID:        GenerateUUID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Items: append(hermioneItems, hermioneItem),
	}
	//hermione.BeforeCreate() // This method will behave as a hook when the DB is implemented.
	characters = append(characters, hermione)

	return characters
}
