package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
)

// The following code fetches data from each museum's API and writes it to the database.
// The archive fetching code is sorted aphabetically by museum name.
//
// 1. Cooper-Hewitt (CooperItem)
// 2. Victoria and Albert Museum (VAItem)

// Fetch the data and return the results to the CooperItems struct.
func (cr CooperItem) Fetch() error {
	log.Info("Fetching items from the Cooper-Hewitt API.")
	var err error
	var resp *http.Response
	var body io.Reader
	var url *url.URL
	var request *http.Request
	var responseBody []byte

	// Build the URL.
	url, err = url.Parse(
		apiBase +
			apiObjectsPath +
			"&access_token=" + os.Getenv("THREADBARE_KEY") +
			"&query=" + sampleQuery,
	)
	if err != nil {
		log.Fatal("Error parsing URL: ", err)
	}

	// Build the request.
	request, err = http.NewRequest("GET", url.String(), body)
	if err != nil {
		log.Error("Error building the request: ", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "threadbare/0.1")

	resp, err = app.Client.Do(request)
	if err != nil {
		log.Error("Error reading HTTP response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"http_error": resp.Status,
			"http_code":  resp.StatusCode,
			"url":        url,
		}).Warn("HTTP error when fetching from API")
	}

	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithField("url", url).WithError(err).Warn("Error reading HTTP response body")
	}

	// Parse the data.
	log.Info("Unmarshalling the data...")

	err = json.Unmarshal(responseBody, &cr)
	if err != nil {
		log.WithFields(log.Fields{
			"url":           url,
			"parsing_error": err,
		}).Error("Error parsing JSON")
	}

	// Check that there is data to write.
	if len(cr.Objects) == 0 {
		log.Warn("No items to save.")
		return nil
	}

	api, err := json.Marshal(cr)
	if err != nil {
		log.Error("Error marshalling API: ", err)
	}

	// Create a timestamp for the items.
	timestamp := "NOW()"
	// Record which museum the items are from.
	archive := "Cooper-Hewitt"

	// Write the data to the database
	for _, item := range cr.Objects {
		_, err := app.DB.Exec(
			`INSERT INTO connthreads.connthreads_items (id, title, date, description, type, medium, url, country, archive, timestamp, api)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			ON CONFLICT (id) DO NOTHING`,
			item.ID,
			item.Title,
			item.Date,
			item.Description,
			item.Type,
			item.Medium,
			item.URL,
			item.Country,
			archive,
			timestamp,
			api,
		)
		if err != nil {
			log.Error("Error writing item to database: ", err)
		}
	}

	return nil
}

// Fetch the data and return the results to the CooperItems struct.
func (cr VAItem) Fetch() error {
	log.Info("Fetching items from the V&A API.")
	var err error
	var resp *http.Response
	var body io.Reader
	var url *url.URL
	var request *http.Request
	var responseBody []byte

	// Build the URL.
	url, err = url.Parse("https://api.vam.ac.uk/v2/objects/search?q=India%20textiles&order_sort=asc&page=1&page_size=100")
	if err != nil {
		log.Fatal("Error parsing URL: ", err)
	}

	// Build the request.
	request, err = http.NewRequest("GET", url.String(), body)
	if err != nil {
		log.Error("Error building the request: ", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "threadbare/0.1")

	resp, err = app.Client.Do(request)
	if err != nil {
		log.Error("Error reading HTTP response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"http_error": resp.Status,
			"http_code":  resp.StatusCode,
			"url":        url,
		}).Warn("HTTP error when fetching from API")
	}

	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithField("url", url).WithError(err).Warn("Error reading HTTP response body")
	}

	// Parse the data.
	log.Info("Unmarshalling the data...")

	err = json.Unmarshal(responseBody, &cr)
	if err != nil {
		log.WithFields(log.Fields{
			"url":           url,
			"parsing_error": err,
		}).Error("Error parsing JSON")
	}

	// Check that there is data to write.
	if len(cr.Records) == 0 {
		log.Warn("No items to save.")
		return nil
	}

	api, err := json.Marshal(cr)
	if err != nil {
		log.Error("Error marshalling API: ", err)
	}

	// Create a timestamp for the items.
	timestamp := "NOW()"
	// Record the archive the items came from.
	archive := "Victoria and Albert Museum"

	// Write the data to the database
	for _, item := range cr.Records {
		_, err := app.DB.Exec(
			`INSERT INTO connthreads.connthreads_items (id, title, date, description, type, medium, url, country, archive, timestamp, api)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			ON CONFLICT (id) DO NOTHING`,
			item.SystemNumber,
			item.PrimaryTitle,
			item.PrimaryDate,
			// item.Description,
			item.Clusters.ObjectType.Terms[0].Value,
			item.Clusters.Material.Terms[0].Value,
			item.Images.IiifImageBaseURL,
			item.Clusters.Place.Terms[0].Value,
			archive,
			timestamp,
			api,
		)
		if err != nil {
			log.Error("Error writing item to database: ", err)
		}
	}

	return nil
}
