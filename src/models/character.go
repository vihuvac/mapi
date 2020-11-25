package models

// Character model.
type Character struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Nickname  string `json:"nickname"`
	Items     []Item `json:"items"`
}

// CharactersResponse for the http request.
type CharactersResponse struct {
	Total      int16       `json:"total"`
	Characters []Character `json:"characters"`
	Link       Link        `json:"link"`
}
