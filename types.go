package main

type Series struct {
	// The url of the series
	BaseURL string `json:"url"`
	// The current episode of the series
	Ep int `json:"ep"`
	// The maximum number of episodes of the series
	Limit int `json:"limit"`
}
