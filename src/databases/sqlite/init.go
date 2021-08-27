package sqlite

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"../../models"
)

// DB reference
var DB *gorm.DB
var err error

// Init the database connection.
func Init() (*gorm.DB, error) {
	log.Println("Initializing SQLite DB...")

	os.Remove("mapi.db")

	DB, err := gorm.Open("sqlite3", "mapi.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer DB.Close()

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
