package models

type SearchResult struct {
	Message  *Message               `json:"message"`
	Summary  *Summary               `json:"summary"` // reserved for future use
	Metadata map[string]interface{} `json:"meta,omitempty"`
	Dist     float64                `json:"dist"`
}

type SearchPayload struct {
	Text     string                 `json:"text"`
	Metadata map[string]interface{} `json:"meta,omitempty"`
}
