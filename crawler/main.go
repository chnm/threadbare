// This program crawls museum APIs to identify textile items and writes the metadata
// to a Postgres database. Currently, it does this with a keyword query of "India textiles"
// in the Cooper Hewitt collection as a proof-of-concept.

package main

import (
	log "github.com/sirupsen/logrus"
)

// Configuration options
const (
	apiBase         = "https://api.collection.cooperhewitt.org/rest/"
	apiObjectsPath  = "?method=cooperhewitt.exhibitions.getObjects"
	apiItemsPerPage = 1000
	apiTimeout      = 60 // timeout limit in seconds
	sampleQuery     = "India%20textiles"
)

var app = &App{}

func main() {
	err := app.Init()
	if err != nil {
		log.Fatal("Error initializing the application: ", err)
	}
	defer app.Shutdown()

	// use fetch.go .Fetch() method to fetch and process data
	results := CooperItem{}
	results.Fetch()

	// Use items.go .Save() method to write to the database
	items := CooperItem{}
	items.Save()

	app.Shutdown() // shutdown the application
	log.Info("Finished the API crawler")
}
