// This program crawls museum APIs to identify textile items and writes the metadata
// to a Postgres database.

package main

import (
	log "github.com/sirupsen/logrus"
)

// Configuration options
const (
	apiItemsPerPage = 1000
	apiTimeout      = 60 // timeout limit in seconds
)

var app = &App{}

func main() {
	err := app.Init()
	if err != nil {
		log.Fatal("Error initializing the application: ", err)
	}
	defer app.Shutdown()

	// use fetch.go .Fetch() method to fetch, process, and write data
	results := CooperItem{}
	err = results.Fetch()
	if err != nil {
		log.Fatal("Error fetching data from Cooper-Hewitt: ", err)
	}

	varesults := VAItem{}
	err = varesults.Fetch()
	if err != nil {
		log.Fatal("Error fetching data from V&A: ", err)
	}

	app.Shutdown() // shutdown the application
	log.Info("Finished the API crawler")
}
