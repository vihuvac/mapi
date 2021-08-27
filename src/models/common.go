package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base structure for all the models.
type Base struct {
	// format: uuid
	// example: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deletedAt"`
}

// BeforeCreate will set a UUID for each new record.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()

	if err != nil {
		return err
	}

	return scope.SetColumn("ID", uuid)
}

// Link model used for paged records.
type Link struct {
	Next string `json:"next"`
}

// CreatedRecordResponse is a generic struct to represent the response of a new record created.
type CreatedRecordResponse struct {
	ID uuid.UUID `json:"id"`
}
