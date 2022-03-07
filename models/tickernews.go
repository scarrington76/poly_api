package model

import "time"

type Tickernews struct {
	Count     int    `json:"count"`
	NextURL   string `json:"next_url"`
	RequestID string `json:"request_id"`
	Results   []struct {
		AmpURL       string    `json:"amp_url"`
		ArticleURL   string    `json:"article_url"`
		Author       string    `json:"author"`
		Description  string    `json:"description"`
		ID           string    `json:"id"`
		ImageURL     string    `json:"image_url"`
		Keywords     []string  `json:"keywords"`
		PublishedUtc time.Time `json:"published_utc"`
		Publisher    struct {
			FaviconURL  string `json:"favicon_url"`
			HomepageURL string `json:"homepage_url"`
			LogoURL     string `json:"logo_url"`
			Name        string `json:"name"`
		} `json:"publisher"`
		Tickers []string `json:"tickers"`
		Title   string   `json:"title"`
	} `json:"results"`
	Status string `json:"status"`
}
