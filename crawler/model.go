package main

// CooperItems describes the collected items from the Cooper Hewitt API.
type CooperItem struct {
	Objects []struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Date        string `json:"date"`
		Description string `json:"description"`
		Type        string `json:"type,omitempty"`
		Medium      string `json:"medium,omitempty"`
		URL         string `json:"url"`
		Country     string `json:"country"`
	} `json:"objects"`
}
