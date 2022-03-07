package model

type Stocksplits struct {
	NextURL string `json:"next_url"`
	Results []struct {
		ExecutionDate string `json:"execution_date"`
		SplitFrom     int    `json:"split_from"`
		SplitTo       int    `json:"split_to"`
		Ticker        string `json:"ticker"`
	} `json:"results"`
	Status string `json:"status"`
}
