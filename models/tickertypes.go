package model

type Tickertypes struct {
	Count     int    `json:"count"`
	RequestID string `json:"request_id"`
	Results   []struct {
		AssetClass  string `json:"asset_class"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Locale      string `json:"locale"`
	} `json:"results"`
	Status string `json:"status"`
}
