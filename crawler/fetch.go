package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
)

// Fetch the data and return the results to the CooperItems struct.
func (cr CooperItem) Fetch() error {
	// query := `
	// INSERT INTO conthreads_items(i d, title, description, url, medium, date, accession_number, department_id, image_url, country, type, items_url, title_raw)
	// VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	// ON CONFLICT DO NOTHING;
	// `

	log.Info("Fetching items from the API.")
	var items []CooperItem
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
	log.Info("Building the request.")
	request, err = http.NewRequest("GET", url.String(), body)
	if err != nil {
		log.Error("Error building the request: ", err)
	}
	log.Info("Setting request headers...")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "threadbare/0.1")

	log.Info("Headers set. Fetching data...")

	// Fetch the data.
	log.Info("Fetching:", cr.Title)

	resp, err = app.Client.Do(request)
	if err != nil {
		fmt.Errorf("Error reading HTTP response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"http_error": resp.Status,
			"http_code":  resp.StatusCode,
			"url":        url,
		}).Warn("HTTP error when fetching from API")
		// return
	}

	defer resp.Body.Close()
	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithField("url", url).WithError(err).Warn("Error reading HTTP response body")
	}

	// Parse the data.
	log.Info("Unmarshalling the data...")

	err = json.Unmarshal(responseBody, &items)
	log.Println(err)
	if err != nil {
		log.WithFields(log.Fields{
			"url":           url,
			"parsing_error": err,
		}).Error("Error parsing JSON")
	}

	log.Debug("Fetched items.")

	// api, err := json.Marshal(ci)
	// if err != nil {
	// 	log.Debug("Error marshalling API: ", err)
	// }

	// stmt, err := app.DB.Prepare(query)
	// if err != nil {
	// 	return err
	// }

	// description := ""
	// if len(ci.Description) > 0 {
	// 	description = ci.Description
	// }

	// // set a timestamp for the item
	// timestamp := "NOW()"

	// _, err = stmt.Exec(ci.ID, ci.Title, description, ci.URL, timestamp, ci.Medium, ci.Date, ci.AccessionNumber, ci.DepartmentID, ci.URL, ci.Country, ci.Type, ci.Images, ci.TitleRaw, api)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// // FetchData sends the API request to the API and returns the results. The results are
// // returned as a CollectionPagination struct.
// func FetchData() ([]CooperItems, error) {
// 	app.Limiters.Collections.Take()

// 	u, _ := url.Parse(
// 		apiBase +
// 			apiObjectsPath +
// 			"&access_token=" + os.Getenv("THREADBARE_KEY") +
// 			"&query=" + sampleQuery,
// 	)

// 	url := u.String()

// 	log.WithField("url", url).Info("Fetching data")
// 	response, err := app.Client.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if response.StatusCode != http.StatusOK {
// 		log.WithFields(log.Fields{
// 			"http_error": response.Status,
// 			"http_code":  response.StatusCode,
// 			"url":        url,
// 		}).Warn("HTTP error when fetching from API")
// 		if response.StatusCode == http.StatusTooManyRequests {
// 			app.Shutdown()
// 			log.Fatal("Quitting, rate limit exceeded.")
// 		}
// 		return nil, fmt.Errorf("HTTP error: %s", response.Status)
// 	}

// 	data, err := io.ReadAll(response.Body)

// 	if err != nil {
// 		return nil, fmt.Errorf("error reading HTTP response body: %w", err)
// 	}

// 	var result ItemList

// 	err = json.Unmarshal(data, &result)
// 	if err != nil {
// 		return nil, fmt.Errorf("error unmarshalling data: %w", err)
// 	}

// 	return result.Results, nil
// }

// CollectionPagination handles pagination objects returned by an API.
// type CollectionPagination struct {
// 	CollectionID string
// 	Pagination   struct {
// 		Current int    `json:"current"`
// 		Next    string `json:"next"`
// 	} `json:"pagination"`
// 	Results []CooperItems `json:"results"`
// 	Title   string        `json:"title"`
// }
