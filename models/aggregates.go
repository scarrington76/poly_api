package model

type Aggregates struct {
	Adjusted   bool   `json:"adjusted"`
	QueryCount int    `json:"queryCount"`
	RequestID  string `json:"request_id"`
	Results    []struct {
		Close                   float64 `json:"c"`
		High                    float64 `json:"h"`
		Low                     float64 `json:"l"`
		Number_transactions     int     `json:"n"`
		Open                    float64 `json:"o"`
		Timestamp               int64   `json:"t"`
		Volume                  int     `json:"v"`
		Volume_weight_avg_price float64 `json:"vw"`
	} `json:"results"`
	ResultsCount int    `json:"resultsCount"`
	Status       string `json:"status"`
	Ticker       string `json:"ticker"`
}
