package model

type Available_exchanges struct {
	Count     int    `json:"count"`
	RequestID string `json:"request_id"`
	Results   []struct {
		Acronym       string `json:"acronym"`
		AssetClass    string `json:"asset_class"`
		ID            int    `json:"id"`
		Locale        string `json:"locale"`
		Mic           string `json:"mic"`
		Name          string `json:"name"`
		OperatingMic  string `json:"operating_mic"`
		ParticipantID string `json:"participant_id"`
		Type          string `json:"type"`
		URL           string `json:"url"`
	} `json:"results"`
	Status string `json:"status"`
}
