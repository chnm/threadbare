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

// Fetch the data and return the results to the CooperItems struct.
func (cr CooperItem) Fetch() error {
	log.Info("Fetching items from the API.")
	var items CooperItem
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

	// var items CooperItem
	err = json.Unmarshal(responseBody, &items)
	if err != nil {
		log.WithFields(log.Fields{
			"url":           url,
			"parsing_error": err,
		}).Error("Error parsing JSON")
	}

	return nil
}
