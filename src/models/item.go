package models

// Item model.
type Item struct {
	Base
	Type            string  `json:"type" validate:"required"`
	Length          float32 `json:"length" validate:"required"`
	Characteristics string  `json:"characteristics" validate:"required"`
	Manufacturer    string  `json:"manufacturer"`
}

// ItemsResponse for the http request.
type ItemsResponse struct {
	Total int16  `json:"total"`
	Items []Item `json:"items"`
	Link  Link   `json:"link"`
}
