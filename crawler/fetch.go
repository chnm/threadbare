package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
)

// FetchData fetches data from the API.
func FetchData() ([]CooperItems, error) {
	app.Limiters.Collections.Take()

	u, _ := url.Parse(
		apiBase +
			apiObjectsPath +
			"&access_token=" + os.Getenv("THREADBARE_KEY") +
			"&query=" + sampleQuery,
	)

	url := u.String()

	log.WithField("url", url).Info("Fetching data")
	response, err := app.Client.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"http_error": response.Status,
			"http_code":  response.StatusCode,
			"url":        url,
		}).Warn("HTTP error when fetching from API")
		if response.StatusCode == http.StatusTooManyRequests {
			app.Shutdown()
			log.Fatal("Quitting, rate limit exceeded.")
		}
		return nil, fmt.Errorf("HTTP error: %s", response.Status)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading HTTP response body: %w", err)
	}

	// unmarshal the data and return it to the caller
	// var result ItemList
	var result CooperItems
	// err = json.Unmarshal(data, &result)
	// if err != nil {
	// return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	// }

	if err := json.Unmarshal(data, &result); err != nil {
		panic(err)
	}

	return result.Results, nil

	// var result ItemList

	// err = json.Unmarshal(data, &result)
	// if err != nil {
	// return nil, fmt.Errorf("error unmarshalling data: %w", err)
	// }

	// return result.Results, nil
}

// CollectionPagination handles pagination objects returned by
// the API.
type CollectionPagination struct {
	CollectionID string
	Pagination   struct {
		Current int    `json:"current"`
		Next    string `json:"next"`
	} `json:"pagination"`
	Results []CooperItems `json:"results"`
	Title   string        `json:"title"`
}
