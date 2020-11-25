package mocks

import (
	"time"

	"../models"
)

// GenerateCharacters generate a list of characters used for testing purposes.
func GenerateCharacters() []models.Character {
	// Init characters var as a slice.
	var characters []models.Character

	// Create Harry's character.
	var harryItems []models.Item
	harry := models.Character{
		FirstName: "Harry",
		LastName:  "Potter",
		Nickname:  "Harry",
		Items: append(harryItems, models.Item{
			models.MaterialInfo{
				Type:            "Wand",
				Length:          11, // Floating number to represent the length in inches.
				Characteristics: "Nice and supple, brother of Voldemort's wand.",
			},
			models.HistoricalInfo{
				Manufacturer: "Garrick Ollivander",
			},
		}),
		Base: models.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	harry.BeforeCreate() // This method will behave as a hook when the DB is implemented.
	characters = append(characters, harry)

	// Create Ron's character.
	var ronItems []models.Item
	ron := models.Character{
		FirstName: "Ronald",
		LastName:  "Weasley",
		Nickname:  "Ron",
		Items: append(ronItems, models.Item{
			models.MaterialInfo{
				Type:            "Wand",
				Length:          12, // Floating number to represent the length in inches.
				Characteristics: "Made of ash wood, and possessed a unicorn hair core.",
			},
			models.HistoricalInfo{
				Manufacturer: "Garrick Ollivander",
			},
		}),
		Base: models.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	ron.BeforeCreate() // This method will behave as a hook when the DB is implemented.
	characters = append(characters, ron)

	// Create Hermione's character.
	var hermioneItems []models.Item
	hermione := models.Character{
		FirstName: "Hermione",
		LastName:  "Granger",
		Nickname:  "Hermione",
		Items: append(hermioneItems, models.Item{
			models.MaterialInfo{
				Type:            "Wand",
				Length:          10.75, // Floating number to represent the length in inches.
				Characteristics: "Made of vine wood, and possessed a dragon heartstring core.",
			},
			models.HistoricalInfo{
				Manufacturer: "Garrick Ollivander",
			},
		}),
		Base: models.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	hermione.BeforeCreate() // This method will behave as a hook when the DB is implemented.
	characters = append(characters, hermione)

	return characters
}
