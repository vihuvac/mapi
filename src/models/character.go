package models

// Character model.
type Character struct {
	Base
	FirstName string `json:"firstName" gorm:"column:firstName" validate:"required"`
	LastName  string `json:"lastName" gorm:"column:lastName" validate:"required"`
	Nickname  string `json:"nickname" gorm:"unique" validate:"required"`
	Items     []Item `json:"items"`
}

// CharactersResponse for the http request.
type CharactersResponse struct {
	Total      int16       `json:"total"`
	Characters []Character `json:"characters"`
	Link       Link        `json:"link"`
}
