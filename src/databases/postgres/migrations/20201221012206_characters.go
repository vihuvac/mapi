package migrations

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20201221012206(db *gorm.DB) {
	msg := "Characters migration 20201221012206 up!"
	fmt.Println(msg)

	err := db.Exec(`
    ALTER TABLE characters ADD CONSTRAINT unique_nickname UNIQUE (nickname);
	`).Error

	if err != nil {
		log.Println("Migration up error:", err)
	}

	fmt.Println(msg + "was successfully executed.")
}

// Down is executed when this migration is rolled back
func Down_20201221012206(db *gorm.DB) {
	msg := "Characters migration 20201221012206 down!"
	fmt.Println(msg)

	err := db.Exec(`
    ALTER TABLE characters DROP CONSTRAINT unique_nickname;
	`).Error

	if err != nil {
		log.Println("Migration down error:", err)
	}

	fmt.Println(msg + "was successfully executed.")
}
