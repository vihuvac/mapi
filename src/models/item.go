package models

// MaterialInfo model.
type MaterialInfo struct {
	Type            string  `json:"type"`
	Length          float32 `json:"length"`
	Characteristics string  `json:"characteristics"`
}

// HistoricalInfo model.
type HistoricalInfo struct {
	Manufacturer string `json:"manufacturer"`
}

// Item model.
type Item struct {
	MaterialInfo
	HistoricalInfo
}
