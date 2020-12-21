package postgres

import (
	"log"
	"os"
	"time"

	"../../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB reference
var DB *gorm.DB
var err error

// Init the database connection.
func Init() (*gorm.DB, error) {
	var postgresURI = os.Getenv("POSTGRES_URI")
	log.Println("POSTGRES_URI:", postgresURI)

	for i := 0; i < 5; i++ {
		DB, err = gorm.Open("postgres", postgresURI) // gorm checks Ping on Open
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return DB, err
	}

	// Create tables if they do not exist.
	if !DB.HasTable(&models.Character{}) {
		log.Println("Creating table &models.Character{}")
		// Gorm creates plural version of Model Character.
		DB.CreateTable(&models.Character{})
	}

	if !DB.HasTable(&models.Item{}) {
		log.Println("Creating table &models.Item{}")
		// Gorm creates plural version of Model Item.
		DB.CreateTable(&models.Item{})
	}

	return DB, err
}
