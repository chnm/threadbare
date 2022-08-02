// This program crawls the Cooper Hewitt API to identify textile items.

package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

func main() {

	req, err := http.NewRequest(
		http.MethodGet,
		apiBase+
			apiObjectsPath+
			"&access_token="+os.Getenv("THREADBARE_KEY")+
			"&query="+sampleQuery,
		nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}
	// Check the response
	log.Println("Response: ", string(response))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
