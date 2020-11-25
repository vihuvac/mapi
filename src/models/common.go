package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Base structure for all the models.
type Base struct {
	// format: uuid
	// example: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// BeforeCreate will set a UUID for each new record.
func (base *Base) BeforeCreate() error {
	uuid, err := uuid.NewV4()

	if err != nil {
		return err
	}

	(*base).ID = uuid

	return nil
}

// Link model used for paged records.
type Link struct {
	Next string `json:"next"`
}
